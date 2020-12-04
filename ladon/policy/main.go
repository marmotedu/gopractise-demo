package main

import "github.com/ory/ladon"

var pol = &ladon.DefaultPolicy{
	// 唯一标识。主要用于数据库检索（必选）。
	ID: "68819e5a-738b-41ec-b03c-b58a1b19d043",

	// 策略描述（可选）。
	Description: "something humanly readable",

	// 主题可以是用户或服务，是“谁能够/不能够对哪些资源做哪些操作”的谁。
	// 可以在< >使用正则表达式。
	Subjects: []string{"max", "peter", "<zac|ken>"},

	// 策略所影响的资源，可以在< >使用正则表达式。
	Resources: []string{
		"myrn:some.domain.com:resource:123", "myrn:some.domain.com:resource:345",
		"myrn:something:foo:<.+>", "myrn:some.domain.com:resource:<(?!protected).*>",
		"myrn:some.domain.com:resource:<[[:digit:]]+>",
	},

	// 策略所影响的操作，可以在< >使用正则表达式。
	Actions: []string{"<create|delete>", "get"},

	// 策略产生的结果是“允许”还是“拒绝”，包括 allow（允许）和 deny（拒绝）。
	// 注意：如果多个策略匹配访问请求，则ladon.DenyAccess将始终覆盖ladon.AllowAccess，并因此拒绝访问。
	Effect: ladon.AllowAccess,

	// 描述策略生效的约束条件。
	Conditions: ladon.Conditions{
		// 在此示例中，仅当所请求的主题也是资源所有者时，该策略才是“活动的”。
		"resourceOwner": &ladon.EqualsSubjectCondition{},

		// 此外，仅当请求的远程IP地址匹配地址范围127.0.0.1/32时，该策略才会匹配。
		"remoteIPAddress": &ladon.CIDRCondition{
			CIDR: "127.0.0.1/32",
		},
	},
}

func main() {}
