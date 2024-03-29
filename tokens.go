package klev

type Token struct {
	TokenID  TokenID   `json:"token_id"`
	Metadata string    `json:"metadata"`
	ACL      []ACLItem `json:"acl"`
	Bearer   string    `json:"bearer,omitempty"`
}

type Tokens struct {
	Tokens []Token `json:"tokens"`
}

type TokenCreateParams struct {
	Metadata string    `json:"metadata,omitempty"`
	ACL      []ACLItem `json:"acl,omitempty"`
}

type TokenUpdateParams struct {
	Metadata *string    `json:"metadata,omitempty"`
	ACL      *[]ACLItem `json:"acl,omitempty"`
}
