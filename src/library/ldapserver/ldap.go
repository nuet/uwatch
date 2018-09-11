// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package ldap provide functions & structure to query a LDAP ldap directory
// For now, it's mainly tested again an MS Active Directory service, see README.md for more information
package ldapserver

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oa/package/ldap"
	"github.com/astaxie/beego"
)

// Basic LDAP authentication service
type Source struct {
	Name             string // canonical name (ie. corporate.ad)
	Host             string // LDAP host
	Port             int    // port number
	UseSSL           bool   // Use SSL
	SkipVerify       bool
	BindDN           string // DN to bind with
	BindPassword     string // Bind DN password
	BaseDN           string // Base search path for users
	UserDN           string // Template for the DN of the user for simple auth
	AttributeName    string // First name attribute
	AttributeSurname string // Surname attribute
	AttributeMail    string // E-mail attribute
	Filter           string // Query filter to validate entry
	titleFilter      string // Query filter to validate entry search for mail
	AdminFilter      string // Query filter to check if user is admin
	Enabled          bool   // if this source is disabled
}

// SaveSource 修改ldap的结构体
type SaveSource struct {
	Add     map[string][]string // 新增属性
	Delete  map[string][]string // 删除属性
	Replace map[string][]string // 修改属性
}

var TestSource = &Source{
	Name:             `ldap.juanpi.org`,
	Host:             `192.168.16.15`,
	Port:             389,
	UseSSL:           false,
	SkipVerify:       false,
	BindDN:           `cn=Manager,dc=juanpi,dc=com`,
	BindPassword:     `123456`,
	BaseDN:           `dc=juanpi,dc=com`,
	UserDN:           "",
	AttributeName:    "givenName",
	AttributeSurname: "displayName",
	AttributeMail:    "mail",
	Filter:           "(&(objectclass=inetorgperson)(cn=%s))",
	titleFilter:      "(&(objectclass=inetorgperson)(title=%s))",
	AdminFilter:      "*",
	Enabled:          true,
}

func NewSource() Source {
	port, _ := beego.AppConfig.Int("ldap_PORT")
	usessl, _ := beego.AppConfig.Bool("ldap_USESSL")
	skipverify, _ := beego.AppConfig.Bool("ldap_SKIPVERIFY")
	return Source{
		Name:             beego.AppConfig.String("ldap_NAME"),
		Host:             beego.AppConfig.String("ldap_HOST"),
		Port:             port,
		UseSSL:           usessl,
		SkipVerify:       skipverify,
		BindDN:           beego.AppConfig.String("ldap_USER"),
		BindPassword:     beego.AppConfig.String("ldap_PASSWD"),
		BaseDN:           beego.AppConfig.String("ldap_BASEDN"),
		UserDN:           "",
		AttributeName:    "givenName",
		AttributeSurname: "displayName",
		AttributeMail:    "mail",
		Filter:           "(&(objectclass=inetorgperson)(cn=%s))",
		titleFilter:      "(&(objectclass=inetorgperson)(title=%s))",
		AdminFilter:      "*",
		Enabled:          true,
	}
}

func ldapDial(ls *Source) (*ldap.Conn, error) {
	ldap.DefaultTimeout = 3 * time.Second
	if ls.UseSSL {
		return ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", ls.Host, ls.Port), &tls.Config{
			InsecureSkipVerify: ls.SkipVerify,
		})
	} else {
		return ldap.Dial("tcp", fmt.Sprintf("%s:%d", ls.Host, ls.Port))
	}
}

// LdapConn ldap连接
func (ls *Source) LdapConn() (l *ldap.Conn, err error) {
	l, err = ldapDial(ls)

	if err != nil {
		ls.Enabled = false
		log.Printf(err.Error())
		return
	}

	err = l.Bind(ls.BindDN, ls.BindPassword)
	if err != nil {
		log.Printf("ERROR: Cannot bind: %s\n", err.Error())
		return
	}
	return l, nil
}

// FindAllGroup 获取所有的组
func (ls *Source) FindAllGroup() []string {
	l, err := ls.LdapConn()
	if err != nil {
		log.Println(err)
	}
	defer l.Close()

	// A search for the user.
	groupFilter := "(objectclass=groupOfUniqueNames)"
	// userFilter := ls.Filter
	fmt.Printf("Searching group filter %s\n", groupFilter)
	search := ldap.NewSearchRequest(
		ls.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		groupFilter,
		[]string{"cn"},
		nil,
	)

	// Ensure we found a user
	sr, err := l.Search(search)
	if err != nil || len(sr.Entries) < 1 {
		fmt.Printf("Failed search using filter[%s]: %v\n", groupFilter, err)
	} else if len(sr.Entries) > 1 {
		sr.PrettyPrint(0)
		fmt.Printf("Filter '%s' returned more than one user.\n", groupFilter)
	}

	groupList := []string{}
	for _, group := range sr.Entries {
		groupList = append(groupList, group.DN)
	}
	return groupList
}

//根据用户花名/拼音 返回所有dn
func (ls *Source) FindAllDNByName(name string) []string {
	var dnString []string
	l, err := ls.LdapConn()
	if err != nil {
		log.Println(err)
		return dnString
	}
	defer l.Close()

	// A search for the user.
	userFilter := fmt.Sprintf(ls.Filter, name)
	// userFilter := ls.Filter
	fmt.Printf("Searching using filter %s\n", userFilter)
	search := ldap.NewSearchRequest(
		ls.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		userFilter,
		[]string{"cn", "username", "mail"},
		nil,
	)

	// Ensure we found a user
	sr, err := l.Search(search)
	if err != nil || len(sr.Entries) < 1 {
		return dnString
	} else if len(sr.Entries) > 1 {
		for i := 0; i < len(sr.Entries); i++ {
			dnString[i] = sr.Entries[i].DN
		}
		return dnString
	}
	dnString[0] = sr.Entries[0].DN
	return dnString
}

// FindUserDN 根据用户花名查找用户
func (ls *Source) FindUserDN(name string) (string, bool, int) {
	l, err := ls.LdapConn()
	if err != nil {
		log.Println(err)
		return "", false, 1000
	}
	defer l.Close()

	// A search for the user.
	userFilter := fmt.Sprintf(ls.Filter, name)
	// userFilter := ls.Filter
	fmt.Printf("Searching using filter %s\n", userFilter)
	search := ldap.NewSearchRequest(
		ls.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		userFilter,
		[]string{"cn", "username", "mail"},
		nil,
	)

	// Ensure we found a user
	sr, err := l.Search(search)
	if err != nil || len(sr.Entries) < 1 {
		//花名登录失败；尝试Title登录
		userTitleFilter := fmt.Sprintf(ls.titleFilter, name)
		fmt.Printf("Searching using titlefilter %s\n", userTitleFilter)
		titlesearch := ldap.NewSearchRequest(
			ls.BaseDN,
			ldap.ScopeWholeSubtree,
			ldap.NeverDerefAliases,
			0,
			0,
			false,
			userTitleFilter,
			[]string{"cn", "username", "mail"},
			nil,
		)
		sr, err1 := l.Search(titlesearch)
		if err1 != nil || len(sr.Entries) < 1 {
			fmt.Printf("Failed search using filter[%s]: %v\n", userTitleFilter, err)
			return "", false, 1002
		} else if len(sr.Entries) > 1 {
			for i := 0; i < len(sr.Entries); i++ {
				array := strings.Split(sr.Entries[i].DN, ",")
				if len(array) == 3 && array[1] == "dc=juanpi" && array[2] == "dc=com" {
					userDN := sr.Entries[i].DN
					fmt.Println("sr.Entries length >1 but dn is", sr.Entries[i].DN)
					return userDN, true, 0
				}
			}
		} else {
			userDN := sr.Entries[0].DN
			if userDN == "" {
				fmt.Printf("LDAP search was succesful, but found no DN!\n")
				return "", false, 1004
			}
			fmt.Printf("userDN '%s' ：\n", userDN)
			return userDN, true, 0
		}
	} else if len(sr.Entries) > 1 {
		for i := 0; i < len(sr.Entries); i++ {
			array := strings.Split(sr.Entries[i].DN, ",")
			if len(array) == 3 && array[1] == "dc=juanpi" && array[2] == "dc=com" {
				userDN := sr.Entries[i].DN
				fmt.Println("sr.Entries length >1 but dn is", sr.Entries[i].DN)
				return userDN, true, 0
			}
		}
	}

	userDN := sr.Entries[0].DN
	if userDN == "" {
		fmt.Printf("LDAP search was succesful, but found no DN!\n")
		return "", false, 1004
	}
	fmt.Printf("userDN '%s' ：\n", userDN)
	return userDN, true, 0
}

// SearchEntry : search an LDAP source if an entry (name, passwd) is valid and in the specific filter
func (ls *Source) SearchEntry(name, passwd string, directBind bool) (string, string, string, bool, int) {
	var userDN string
	var errcode int
	if directBind {
		fmt.Printf("LDAP will bind directly via UserDN template: %s\n", ls.UserDN)
		userDN = fmt.Sprintf(ls.UserDN, name)
	} else {
		fmt.Printf("LDAP will use BindDN.\n")

		var found bool
		//"" false 1003
		userDN, found, errcode = ls.FindUserDN(name)
		if !found {
			return "", "", "", false, errcode
		}
	}

	l, err := ls.LdapConn()
	if err != nil {
		return "", "", "", false, 1000
	}
	defer l.Close()

	fmt.Println("Binding with userDN: %s", userDN)
	err = l.Bind(userDN, passwd)
	if err != nil {
		fmt.Printf("LDAP auth. failed for %s, reason: %v\n", userDN, err)
		return "", "", "", false, 2000
	}

	fmt.Printf("Bound successfully with userDN: %s\n", userDN)

	userFilter := fmt.Sprintf(ls.Filter, name)
	search := ldap.NewSearchRequest(
		userDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, userFilter,
		[]string{ls.AttributeName, ls.AttributeSurname, ls.AttributeMail},
		nil)

	sr, err := l.Search(search)
	if err != nil {
		fmt.Printf("LDAP Search failed unexpectedly! (%v)\n", err)
		return "", "", "", false, 2001
	} else if len(sr.Entries) < 1 {
		//尝试用title来登录
		userTitleFilter := fmt.Sprintf(ls.titleFilter, name)
		titlesearch := ldap.NewSearchRequest(
			userDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, userTitleFilter,
			[]string{ls.AttributeName, ls.AttributeSurname, ls.AttributeMail},
			nil)

		sr, err1 := l.Search(titlesearch)

		if err1 != nil {
			fmt.Printf("LDAP Search failed unexpectedly! (%v)\n", err)
			return "", "", "", false, 2001
		} else if len(sr.Entries) < 1 {
			if directBind {
				fmt.Printf("User filter inhibited user login.\n")
			} else {
				fmt.Printf("LDAP Search failed unexpectedly! (0 entries)\n")
			}

			return "", "", "", false, 2002
		} else {
			name_attr := sr.Entries[0].GetAttributeValue(ls.AttributeName)
			sn_attr := sr.Entries[0].GetAttributeValue(ls.AttributeSurname)
			mail_attr := sr.Entries[0].GetAttributeValue(ls.AttributeMail)

			return name_attr, sn_attr, mail_attr, true, 0
		}
	}
	name_attr := sr.Entries[0].GetAttributeValue(ls.AttributeName)
	sn_attr := sr.Entries[0].GetAttributeValue(ls.AttributeSurname)
	mail_attr := sr.Entries[0].GetAttributeValue(ls.AttributeMail)

	return name_attr, sn_attr, mail_attr, true, 0
}

// Add 添加数据 增加用户
func (ls *Source) Add(data map[string][]string, dn string) (err error) {
	l, err := ls.LdapConn()
	if err != nil {
		log.Printf("ERROR: Conn Ldap fail: %s\n", err.Error())
		return
	}
	defer l.Close()
	//dn := fmt.Sprintf("cn=%s,%s",data["cn"][0],ls.BaseDN)
	addRes := ldap.NewAddRequest(dn)
	for key, value := range data {
		addRes.Attribute(key, value)
	}

	err = l.Add(addRes)
	if err != nil {
		log.Printf("ERROR: Add %s fail: %s\n", dn, err.Error())
		return
	}
	return
}

// Delete 数据删除
func (ls *Source) Delete(dn string) (err error) {
	l, err := ls.LdapConn()
	if err != nil {
		return
	}
	defer l.Close()

	delRequest := ldap.NewDelRequest(dn, nil)
	err = l.Del(delRequest)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Save LDAP数据更新保存
func (ls *Source) Save(dn string, attributes SaveSource) (err error) {
	modifyRequest := ldap.NewModifyRequest(dn)
	for attrType, attrVals := range attributes.Add {
		modifyRequest.Add(attrType, attrVals)
	}
	for attrType, attrVals := range attributes.Delete {
		modifyRequest.Delete(attrType, attrVals)
	}
	for attrType, attrVals := range attributes.Replace {
		modifyRequest.Replace(attrType, attrVals)
	}

	l, err := ls.LdapConn()
	if err != nil {
		return
	}
	defer l.Close()

	err = l.Modify(modifyRequest)

	return
}

// GetGroupMember 获取组员dn列表
func (ls *Source) GetGroupMember(cn string) (res []string, err error) {
	l, err := ls.LdapConn()
	if err != nil {
		return
	}
	defer l.Close()
	filter := fmt.Sprintf("(&(%s))", cn)
	var Attributes []string = []string{}
	search := ldap.NewSearchRequest(
		ls.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,
		Attributes,
		nil)
	sr, err := l.Search(search)
	if err != nil {
		log.Println("ERROR: %s\n", err.Error())
		return
	}
	if len(sr.Entries) > 0 {
		res = sr.Entries[0].GetAttributeValues("uniqueMember")
	}
	return
}

// ModifyGroupUser 修改组用户
func (ls *Source) ModifyGroupUser(groupsCn, addUsersDn, delUsersDn []string) error {
	msg := ""
	for _, groupCn := range groupsCn {
		groupDn := groupCn + "," + ls.BaseDN
		attr := SaveSource{}
		if len(addUsersDn) > 0 {
			attr.Add = map[string][]string{"uniqueMember": addUsersDn}
		}
		if len(delUsersDn) > 0 {
			attr.Delete = map[string][]string{"uniqueMember": delUsersDn}
		}

		err := ls.Save(groupDn, attr)
		if err != nil {
			msg += fmt.Sprintf("修改组用户%s失败:%s；", groupCn, err.Error())
		}
	}
	if msg != "" {
		err := errors.New(msg)
		return err
	}
	return nil
}
