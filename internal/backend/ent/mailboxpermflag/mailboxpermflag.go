// Code generated by ent, DO NOT EDIT.

package mailboxpermflag

const (
	// Label holds the string label denoting the mailboxpermflag type in the database.
	Label = "mailbox_perm_flag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// Table holds the table name of the mailboxpermflag in the database.
	Table = "mailbox_perm_flags"
)

// Columns holds all SQL columns for mailboxpermflag fields.
var Columns = []string{
	FieldID,
	FieldValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "mailbox_perm_flags"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"mailbox_permanent_flags",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
