package comments

func CodifyText(text string) string {
	return "```\n" + text + "\n```\n"
}

func AttachComment(text string, comment string) string {
	return comment + "\n" + text + "\n"
}

func MakeComment(text string, comment string, codify bool) string {
	if codify == true {
		return AttachComment(CodifyText(text), comment)
	}
	return AttachComment(text, comment)
}

func MakeSingleComment(rawText []string, rawComment []string, codify bool) string {
	if len(rawText) == 1 {
		if len(rawComment) > 0 {
			return MakeComment(rawText[0], rawComment[0], codify)
		}
		return MakeComment(rawText[0], "", codify)
	}
	var result string
	for i, t := range rawText {
		if i >= len(rawComment) {
			result += MakeComment(t, "", codify)
		} else {
			result += MakeComment(t, rawComment[i], codify)
		}
	}
	return result
}
