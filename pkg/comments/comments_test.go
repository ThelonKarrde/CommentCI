package comments_test

import (
	"github.com/ThelonKarrde/CommentCI/pkg/comments"
	"testing"
)

func TestCodifyTest(t *testing.T) {
	testString := "comment"
	rString := "```\ncomment\n```\n"
	codifiedString := comments.CodifyText(testString)
	if codifiedString != rString {
		t.Error("Fail to codify string! " + codifiedString)
	}
}

func TestAttachComment(t *testing.T) {
	text := "text"
	comment := "comment"
	rString := "comment\ntext\n"
	cmtString := comments.AttachComment(text, comment)
	if cmtString != rString {
		t.Error("Fail to attach comment! " + cmtString)
	}
}

func TestMakeCommentNoCodify(t *testing.T) {
	text := "text"
	comment := "comment"
	rString := "comment\ntext\n"
	cmtString := comments.MakeComment(text, comment, false)
	if cmtString != rString {
		t.Error("Fail to make a comment without codify! " + cmtString)
	}
}

func TestMakeCommentWithCodify(t *testing.T) {
	text := "text"
	comment := "comment"
	rString := "comment\n```\ntext\n```\n\n"
	cmtString := comments.MakeComment(text, comment, true)
	if cmtString != rString {
		t.Error("Fail to make a comment with codify! Real: " + cmtString + " Desired: " + rString)
	}
}

func TestMakeSingleCommentMText(t *testing.T) {
	texts := []string{"text1", "text2"}
	cmts := []string{"comment1", "comment2"}
	rString := "comment1\ntext1\ncomment2\ntext2\n"

	cmtString := comments.MakeSingleComment(texts, cmts, false)
	if cmtString != rString {
		t.Error("Fail to make a single comment with codify! Real: " + cmtString + " Desired: " + rString)
	}
}

func TestMakeSingleCommentSText(t *testing.T) {
	texts := []string{"text1"}
	cmts := []string{"comment1"}
	rString := "comment1\ntext1\n"

	cmtString := comments.MakeSingleComment(texts, cmts, false)
	if cmtString != rString {
		t.Error("Fail to make a single comment with codify! Real: " + cmtString + " Desired: " + rString)
	}
}

func TestMakeSingleCommentSTextNoComment(t *testing.T) {
	texts := []string{"text1"}
	var cmts []string
	rString := "\ntext1\n"

	cmtString := comments.MakeSingleComment(texts, cmts, false)
	if cmtString != rString {
		t.Error("Fail to make a single comment without codify and no comments! Real: " + cmtString + " Desired: " + rString)
	}
}

func TestMakeSingleCommentMTextNoComment(t *testing.T) {
	texts := []string{"text1", "text2"}
	var cmts []string
	rString := "\ntext1\n\ntext2\n"

	cmtString := comments.MakeSingleComment(texts, cmts, false)
	if cmtString != rString {
		t.Error("Fail to make a single comment without codify and no comments! Real: " + cmtString + " Desired: " + rString)
	}
}

func TestMakeSingleCommentMTextSComments(t *testing.T) {
	texts := []string{"text1", "text2"}
	cmts := []string{"comment1"}
	rString := "comment1\ntext1\n\ntext2\n"

	cmtString := comments.MakeSingleComment(texts, cmts, false)
	if cmtString != rString {
		t.Error("Fail to make a single comment without codify and 2/1 text comments! Real: " + cmtString + " Desired: " + rString)
	}
}
