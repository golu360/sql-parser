package states

type State int8

const (
	TOKENIZE                  State = 0
	TOKENIZE_COMPLETE         State = 1
	QUERY_TYPE_PARSE          State = 2
	QUERY_TYPE_PARSE_COMPLETE State = 3
	SELECT_FIELD_PARSE        State = 4
)
