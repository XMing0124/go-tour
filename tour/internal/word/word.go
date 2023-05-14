package word

import (
	"strings"
	"unicode"
)

// ToUpper 单词全部转换为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 单词全部转换为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// UnderLineToUpperCamelCase converts underline to upper camel
// user_name -> UserName
func UnderlineToUpperCamelCase(s string) string {
	s = strings.Replace(s,"_"," ",-1)
	s = strings.Title(s)
	return strings.Replace(s," ","",-1)
}

// UnderLineToLowerCamelCase converts underline to lower camel
// user_name -> userName
func UnderLineToLowerCamelCase(s string) string {
	s = UnderlineToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToUnderline converts camel to underline
// userName -> user_name
func CamelCaseToUnderline(s string) string {
	var output []rune
	for i, r := range s {
		if i==0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
            output = append(output, '_')
        }
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}