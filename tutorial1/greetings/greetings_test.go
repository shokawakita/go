package greetings 

import ( 
    "testing" 
    "regexp" 
) 

// TestHelloName は名前で greetings.Hello を呼び出し、
// 戻り値の有効性を確認します。
func TestHelloName(t *testing.T) { 
    name := "Gladys" 
    want := regexp.MustCompile(`\b`+name+`\b`) 
    msg, err := Hello("Gladys") 
    if !want.MatchString(msg) || err != nil { 
        t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want) 
    } 
} 

// TestHelloEmpty は空の文字列で greetings.Hello を呼び出し、
// エラーを確認します。
func TestHelloEmpty(t *testing.T) { 
    msg, err := Hello("") 
    if msg != "" || err == nil { 
        t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err) 
    } 
}