// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package gofly

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonD2b7633eDecodeGithubComNoisyscannerGoflyGofly(in *jlexer.Lexer, out *VerbContainer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				in.Delim('[')
				if out.Data == nil {
					if !in.IsDelim(']') {
						out.Data = make([]Verb, 0, 1)
					} else {
						out.Data = []Verb{}
					}
				} else {
					out.Data = (out.Data)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Verb
					(v1).UnmarshalEasyJSON(in)
					out.Data = append(out.Data, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComNoisyscannerGoflyGofly(out *jwriter.Writer, in VerbContainer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix[1:])
		if in.Data == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Data {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v VerbContainer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComNoisyscannerGoflyGofly(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v VerbContainer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComNoisyscannerGoflyGofly(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *VerbContainer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComNoisyscannerGoflyGofly(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *VerbContainer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComNoisyscannerGoflyGofly(l, v)
}
func easyjsonD2b7633eDecodeGithubComNoisyscannerGoflyGofly1(in *jlexer.Lexer, out *Verb) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int(in.Int())
		case "i":
			out.Infinitive = string(in.String())
		case "ni":
			out.NormalisedInfinitive = string(in.String())
		case "e":
			out.English = string(in.String())
		case "hid":
			out.HelperID = int(in.Int())
		case "ih":
			out.IsHelper = bool(in.Bool())
		case "ir":
			out.IsReflexive = bool(in.Bool())
		case "conjugations":
			(out.Conjugations).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComNoisyscannerGoflyGofly1(out *jwriter.Writer, in Verb) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"i\":"
		out.RawString(prefix)
		out.String(string(in.Infinitive))
	}
	if in.NormalisedInfinitive != "" {
		const prefix string = ",\"ni\":"
		out.RawString(prefix)
		out.String(string(in.NormalisedInfinitive))
	}
	{
		const prefix string = ",\"e\":"
		out.RawString(prefix)
		out.String(string(in.English))
	}
	if in.HelperID != 0 {
		const prefix string = ",\"hid\":"
		out.RawString(prefix)
		out.Int(int(in.HelperID))
	}
	{
		const prefix string = ",\"ih\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsHelper))
	}
	{
		const prefix string = ",\"ir\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsReflexive))
	}
	{
		const prefix string = ",\"conjugations\":"
		out.RawString(prefix)
		(in.Conjugations).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Verb) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComNoisyscannerGoflyGofly1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Verb) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComNoisyscannerGoflyGofly1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Verb) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComNoisyscannerGoflyGofly1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Verb) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComNoisyscannerGoflyGofly1(l, v)
}
func easyjsonD2b7633eDecodeGithubComNoisyscannerGoflyGofly2(in *jlexer.Lexer, out *Tense) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int(in.Int())
		case "identifier":
			out.Identifier = string(in.String())
		case "displayName":
			out.DisplayName = string(in.String())
		case "order":
			out.Order = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComNoisyscannerGoflyGofly2(out *jwriter.Writer, in Tense) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"identifier\":"
		out.RawString(prefix)
		out.String(string(in.Identifier))
	}
	{
		const prefix string = ",\"displayName\":"
		out.RawString(prefix)
		out.String(string(in.DisplayName))
	}
	{
		const prefix string = ",\"order\":"
		out.RawString(prefix)
		out.Int(int(in.Order))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Tense) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComNoisyscannerGoflyGofly2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Tense) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComNoisyscannerGoflyGofly2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Tense) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComNoisyscannerGoflyGofly2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Tense) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComNoisyscannerGoflyGofly2(l, v)
}
func easyjsonD2b7633eDecodeGithubComNoisyscannerGoflyGofly3(in *jlexer.Lexer, out *Pronoun) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int(in.Int())
		case "identifier":
			out.Identifier = string(in.String())
		case "displayName":
			out.DisplayName = string(in.String())
		case "reflexive":
			out.Reflexive = string(in.String())
		case "order":
			out.Order = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComNoisyscannerGoflyGofly3(out *jwriter.Writer, in Pronoun) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"identifier\":"
		out.RawString(prefix)
		out.String(string(in.Identifier))
	}
	{
		const prefix string = ",\"displayName\":"
		out.RawString(prefix)
		out.String(string(in.DisplayName))
	}
	{
		const prefix string = ",\"reflexive\":"
		out.RawString(prefix)
		out.String(string(in.Reflexive))
	}
	{
		const prefix string = ",\"order\":"
		out.RawString(prefix)
		out.Int(int(in.Order))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Pronoun) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComNoisyscannerGoflyGofly3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Pronoun) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComNoisyscannerGoflyGofly3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Pronoun) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComNoisyscannerGoflyGofly3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Pronoun) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComNoisyscannerGoflyGofly3(l, v)
}
func easyjsonD2b7633eDecodeGithubComNoisyscannerGoflyGofly4(in *jlexer.Lexer, out *Language) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int(in.Int())
		case "lang":
			out.Lang = string(in.String())
		case "code":
			out.Code = string(in.String())
		case "locale":
			out.Locale = string(in.String())
		case "version":
			out.Version = int(in.Int())
		case "schemaVersion":
			out.SchemaVersion = int(in.Int())
		case "hasReflexives":
			out.HasReflexives = bool(in.Bool())
		case "hasHelpers":
			out.HasHelpers = bool(in.Bool())
		case "tenses":
			easyjsonD2b7633eDecode(in, &out.Tenses)
		case "pronouns":
			easyjsonD2b7633eDecode1(in, &out.Pronouns)
		case "verbs":
			easyjsonD2b7633eDecode2(in, &out.Verbs)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComNoisyscannerGoflyGofly4(out *jwriter.Writer, in Language) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"lang\":"
		out.RawString(prefix)
		out.String(string(in.Lang))
	}
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix)
		out.String(string(in.Code))
	}
	{
		const prefix string = ",\"locale\":"
		out.RawString(prefix)
		out.String(string(in.Locale))
	}
	{
		const prefix string = ",\"version\":"
		out.RawString(prefix)
		out.Int(int(in.Version))
	}
	{
		const prefix string = ",\"schemaVersion\":"
		out.RawString(prefix)
		out.Int(int(in.SchemaVersion))
	}
	{
		const prefix string = ",\"hasReflexives\":"
		out.RawString(prefix)
		out.Bool(bool(in.HasReflexives))
	}
	{
		const prefix string = ",\"hasHelpers\":"
		out.RawString(prefix)
		out.Bool(bool(in.HasHelpers))
	}
	{
		const prefix string = ",\"tenses\":"
		out.RawString(prefix)
		easyjsonD2b7633eEncode(out, in.Tenses)
	}
	{
		const prefix string = ",\"pronouns\":"
		out.RawString(prefix)
		easyjsonD2b7633eEncode1(out, in.Pronouns)
	}
	{
		const prefix string = ",\"verbs\":"
		out.RawString(prefix)
		easyjsonD2b7633eEncode2(out, in.Verbs)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Language) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComNoisyscannerGoflyGofly4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Language) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComNoisyscannerGoflyGofly4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Language) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComNoisyscannerGoflyGofly4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Language) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComNoisyscannerGoflyGofly4(l, v)
}
func easyjsonD2b7633eDecode2(in *jlexer.Lexer, out *struct{ Data []Verb }) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				in.Delim('[')
				if out.Data == nil {
					if !in.IsDelim(']') {
						out.Data = make([]Verb, 0, 1)
					} else {
						out.Data = []Verb{}
					}
				} else {
					out.Data = (out.Data)[:0]
				}
				for !in.IsDelim(']') {
					var v4 Verb
					(v4).UnmarshalEasyJSON(in)
					out.Data = append(out.Data, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncode2(out *jwriter.Writer, in struct{ Data []Verb }) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix[1:])
		if in.Data == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Data {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonD2b7633eDecode1(in *jlexer.Lexer, out *struct{ Data []Pronoun }) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				in.Delim('[')
				if out.Data == nil {
					if !in.IsDelim(']') {
						out.Data = make([]Pronoun, 0, 1)
					} else {
						out.Data = []Pronoun{}
					}
				} else {
					out.Data = (out.Data)[:0]
				}
				for !in.IsDelim(']') {
					var v7 Pronoun
					(v7).UnmarshalEasyJSON(in)
					out.Data = append(out.Data, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncode1(out *jwriter.Writer, in struct{ Data []Pronoun }) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix[1:])
		if in.Data == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Data {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonD2b7633eDecode(in *jlexer.Lexer, out *struct{ Data []Tense }) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				in.Delim('[')
				if out.Data == nil {
					if !in.IsDelim(']') {
						out.Data = make([]Tense, 0, 1)
					} else {
						out.Data = []Tense{}
					}
				} else {
					out.Data = (out.Data)[:0]
				}
				for !in.IsDelim(']') {
					var v10 Tense
					(v10).UnmarshalEasyJSON(in)
					out.Data = append(out.Data, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncode(out *jwriter.Writer, in struct{ Data []Tense }) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix[1:])
		if in.Data == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Data {
				if v11 > 0 {
					out.RawByte(',')
				}
				(v12).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonD2b7633eDecodeGithubComNoisyscannerGoflyGofly5(in *jlexer.Lexer, out *ConjugationContainer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				in.Delim('[')
				if out.Data == nil {
					if !in.IsDelim(']') {
						out.Data = make([]Conjugation, 0, 1)
					} else {
						out.Data = []Conjugation{}
					}
				} else {
					out.Data = (out.Data)[:0]
				}
				for !in.IsDelim(']') {
					var v13 Conjugation
					(v13).UnmarshalEasyJSON(in)
					out.Data = append(out.Data, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComNoisyscannerGoflyGofly5(out *jwriter.Writer, in ConjugationContainer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix[1:])
		if in.Data == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.Data {
				if v14 > 0 {
					out.RawByte(',')
				}
				(v15).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ConjugationContainer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComNoisyscannerGoflyGofly5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ConjugationContainer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComNoisyscannerGoflyGofly5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ConjugationContainer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComNoisyscannerGoflyGofly5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ConjugationContainer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComNoisyscannerGoflyGofly5(l, v)
}
func easyjsonD2b7633eDecodeGithubComNoisyscannerGoflyGofly6(in *jlexer.Lexer, out *Conjugation) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "c":
			out.Conjugation = string(in.String())
		case "nc":
			out.NormalisedConjugation = string(in.String())
		case "tid":
			out.TenseID = int(in.Int())
		case "pid":
			out.PronounID = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComNoisyscannerGoflyGofly6(out *jwriter.Writer, in Conjugation) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"c\":"
		out.RawString(prefix[1:])
		out.String(string(in.Conjugation))
	}
	if in.NormalisedConjugation != "" {
		const prefix string = ",\"nc\":"
		out.RawString(prefix)
		out.String(string(in.NormalisedConjugation))
	}
	{
		const prefix string = ",\"tid\":"
		out.RawString(prefix)
		out.Int(int(in.TenseID))
	}
	{
		const prefix string = ",\"pid\":"
		out.RawString(prefix)
		out.Int(int(in.PronounID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Conjugation) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComNoisyscannerGoflyGofly6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Conjugation) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComNoisyscannerGoflyGofly6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Conjugation) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComNoisyscannerGoflyGofly6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Conjugation) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComNoisyscannerGoflyGofly6(l, v)
}
