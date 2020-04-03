package main

type Report struct {
	ToolsVersion string
	Information  Information
	IssueTypes   []IssueType `xml:"IssueTypes>IssueType"`
	Issues       []Project   `xml:"Issues>Project"`
}

type Information struct {
	Solution        string
	InspectionScope string `xml:">Element"`
}

type IssueType struct {
	Id          string `xml:",attr"`
	Category    string `xml:",attr"`
	CategoryId  string `xml:",attr"`
	SubCategory string `xml:",attr"`
	Description string `xml:",attr"`
	Severity    string `xml:",attr"`
	Global      string `xml:",attr"`
	WikiUrl     string `xml:",attr"`
}

type Project struct {
	Name   string         `xml:",attr"`
	Issues []ProjectIssue `xml:"Issue"`
}

type ProjectIssue struct {
	TypeId  string `xml:",attr"`
	File    string `xml:",attr"`
	Offset  string `xml:",attr"`
	Line    string `xml:",attr"`
	Message string `xml:",attr"`
}
