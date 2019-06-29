package main

import (
	"bytes"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// var out io.Writer = os.Stdout

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	want := strconv.Quote("a:2\nb:1\n")
	got := strconv.Quote(buf.String())

	if got != want {
		t.Errorf("actual: %s does not match expected: %s", got, want)
	}
}

var sortLettersTestData = []struct {
	name  string
	input string
	want  string
}{
	{name: "lab example", input: "aba", want: "a:2,b:1"},
	{name: "trailing newline", input: "aba\n", want: "\n:1,a:2,b:1"},
	{name: "duplicate entries", input: "\t\t\\aba\n", want: "\t:2,\n:1,\\:1,a:2,b:1"},
	{name: "one letter", input: "a", want: "a:1"},
	{name: "one repeated letter", input: "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", want: "b:61"},
	{name: "repeated entry", input: "abababababababab", want: "a:8,b:8"},
	{name: "empty string", input: "", want: ""},
	{name: "random example", input: "blahflehblahbleh", want: "a:2,b:3,e:2,f:1,h:4,l:4"},
	{name: "whitespace", input: "\t\t   \t\t\t \t\n\n", want: "\t:6,\n:2, :4"},
	{name: "Katakana",
		input: "  イロハニホヘト チリヌルヲ ワカヨタレソ ツネナラム ツネナラム ウヰノオクヤマ ケフコエテ アサキユメミシ ヱヒモセスン",
		want: " :10,ア:1,イ:1,ウ:1,エ:1,オ:1,カ:1,キ:1,ク:1,ケ:1,コ:1,サ:1,シ:1,ス:1,セ:1,ソ:1,タ:1,チ:1,ツ:2,テ:1,ト:1,ナ:2," +
			"ニ:1,ヌ:1,ネ:2,ノ:1,ハ:1,ヒ:1,フ:1,ヘ:1,ホ:1,マ:1,ミ:1,ム:2,メ:1,モ:1,ヤ:1,ユ:1,ヨ:1,ラ:2,リ:1,ル:1,レ:1,ロ:1,ワ:1,ヰ:1," +
			"ヱ:1,ヲ:1,ン:1"},
	{name: "Polish",
		input: "Pchnąć w tę łódź jeża lub ośm skrzyń fig",
		want: " :8,P:1,a:1,b:1,c:1,d:1,e:1,f:1,g:1,h:1,i:1,j:1,k:1,l:1,m:1,n:1,o:1,r:1,s:1,t:1,u:1,w:1,y:1,z:1,ó:1,ą:1,ć:1," +
			"ę:1,ł:1,ń:1,ś:1,ź:1,ż:1"},
	{name: "thai",
		input: "  ๏ เป็นมนุษย์สุดประเสริฐเลิศคุณค่า กว่าบรรดาฝูงสัตว์เดรัจฉาน " +
			"จงฝ่าฟันพัฒนาวิชาการ อย่าล้างผลาญฤๅเข่นฆ่าบีฑาใคร " +
			"ม่ถือโทษโกรธแช่งซัดฮึดฮัดด่า     หัดอภัยเหมือนกีฬาอัชฌาสัย	" +
			"ปฏิบัติประพฤติกฎกำหนดใจ พูดจาให้จ๊ะๆ จ๋าๆ น่าฟังเอย ฯ",
		want: "\t:1, :16,ก:6,ข:1,ค:3,ฆ:1,ง:5,จ:6,ฉ:1,ช:3,ซ:1,ฌ:1,ญ:1,ฎ:1,ฏ:1,ฐ:1,ฑ:1,ฒ:1,ณ:1,ด:10,ต:3,ถ:1," +
			"ท:1,ธ:1,น:9,บ:3,ป:4,ผ:1,ฝ:2,พ:3,ฟ:2,ภ:1,ม:3,ย:5,ร:9,ฤ:2,ล:3,ว:3,ศ:1,ษ:2,ส:4,ห:4,ฬ:1,อ:6,ฮ:2,ฯ:1,ะ:3,ั" +
			":12,า:19,ำ:1,ิ:6,ี:2,ึ:1,ื:2,ุ:3,ู:2,เ:7,แ:1,โ:2,ใ:3,ๅ:1,ๆ:2,็:1,่:10,้:2,๊:1,๋:1,์:2,๏:1"},
	{name: "Russian",
		input: "Съешь же ещё этих мягких французских булок да выпей чаю",
		want: " :9,С:1,а:3,б:1,в:1,г:1,д:1,е:4,ж:1,з:1,и:3,й:1,к:3,л:1,м:1,н:1,о:1,п:1,р:1," +
			"с:1,т:1,у:2,ф:1,х:3,ц:1,ч:1,ш:1,щ:1,ъ:1,ы:1,ь:1,э:1,ю:1,я:1,ё:1"},
	{name: "Emojis",
		input: "👾 🙇 💁 🙅 🙆 🙋 🙎 🐵 🙈 🙉 🙊 ❤️ 💔 💌 💕 💞 💓 💗 💖 💘 💝 💟 💜 💛 💚 💙",
		want: " :25,❤:1,️:1,🐵:1,👾:1,💁:1,💌:1,💓:1,💔:1,💕:1,💖:1,💗:1,💘:1,💙:1,💚:1,💛:1,💜:1,💝:1,💞:1,💟:1," +
			"🙅:1,🙆:1,🙇:1,🙈:1,🙉:1,🙊:1,🙋:1,🙎:1"},
}

func TestSortLetters(t *testing.T) {
	for _, tt := range sortLettersTestData {
		tt := tt // as per sean- suggestion on this discussion https://github.com/kyoh86/scopelint/issues/4
		t.Run(tt.name, func(t *testing.T) {
			got := strings.Join(sortLetters(letters(tt.input)), ",")
			if got != tt.want {
				t.Errorf("sortLetters() = \n%v,\nwant \n%v", got, tt.want)
			}
		})
	}
}

var lettersTestData = []struct {
	name  string
	input string
	want  map[rune]int
}{
	{name: "lab example", input: "aba", want: map[rune]int{97: 2, 98: 1}},
	{name: "trailing newline", input: "aba\n", want: map[rune]int{10: 1, 97: 2, 98: 1}},
	{name: "duplicate entries", input: "\t\t\\aba\n", want: map[rune]int{9: 2, 10: 1, 92: 1, 97: 2, 98: 1}},
	{name: "one letter", input: "a", want: map[rune]int{97: 1}},
	{name: "one repeated letter",
		input: "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
		want:  map[rune]int{98: 61}},
	{name: "repeated entry", input: "abababababababab", want: map[rune]int{97: 8, 98: 8}},
	{name: "empty string", input: "", want: map[rune]int{}},
	{name: "random example", input: "blahflehblahbleh", want: map[rune]int{97: 2, 98: 3, 101: 2, 102: 1, 104: 4, 108: 4}},
	{name: "whitespace", input: "\t\t   \t\t\t \t\n\n", want: map[rune]int{9: 6, 10: 2, 32: 4}},
	{name: "Katakana",
		input: "  イロハニホヘト チリヌルヲ ワカヨタレソ ツネナラム ツネナラム ウヰノオクヤマ ケフコエテ アサキユメミシ ヱヒモセスン",
		want: map[rune]int{32: 10, 12450: 1, 12452: 1, 12454: 1, 12456: 1, 12458: 1, 12459: 1, 12461: 1,
			12463: 1, 12465: 1, 12467: 1, 12469: 1, 12471: 1, 12473: 1, 12475: 1, 12477: 1, 12479: 1, 12481: 1,
			12484: 2, 12486: 1, 12488: 1, 12490: 2, 12491: 1, 12492: 1, 12493: 2, 12494: 1, 12495: 1, 12498: 1,
			12501: 1, 12504: 1, 12507: 1, 12510: 1, 12511: 1, 12512: 2, 12513: 1, 12514: 1, 12516: 1, 12518: 1,
			12520: 1, 12521: 2, 12522: 1, 12523: 1, 12524: 1, 12525: 1, 12527: 1, 12528: 1, 12529: 1, 12530: 1,
			12531: 1}},
	{name: "Polish",
		input: "Pchnąć w tę łódź jeża lub ośm skrzyń fig",
		want: map[rune]int{32: 8, 80: 1, 97: 1, 98: 1, 99: 1, 100: 1, 101: 1, 102: 1, 103: 1, 104: 1, 105: 1, 106: 1,
			107: 1, 108: 1, 109: 1, 110: 1, 111: 1, 114: 1, 115: 1, 116: 1, 117: 1, 119: 1, 121: 1, 122: 1, 243: 1, 261: 1,
			263: 1, 281: 1, 322: 1, 324: 1, 347: 1, 378: 1, 380: 1}},
	{name: "thai",
		input: "  ๏ เป็นมนุษย์สุดประเสริฐเลิศคุณค่า กว่าบรรดาฝูงสัตว์เดรัจฉาน จงฝ่าฟันพัฒนาวิชาการ" +
			" อย่าล้างผลาญฤๅเข่นฆ่าบีฑาใคร ม่ถือโทษโกรธแช่งซัดฮึดฮัดด่า" +
			"     หัดอภัยเหมือนกีฬาอัชฌาสัย	ปฏิบัติประพฤติกฎกำหนดใจ พูดจาให้จ๊ะๆ จ๋าๆ น่าฟังเอย ฯ",
		want: map[rune]int{9: 1, 32: 16, 3585: 6, 3586: 1, 3588: 3, 3590: 1, 3591: 5, 3592: 6, 3593: 1, 3594: 3,
			3595: 1, 3596: 1, 3597: 1, 3598: 1, 3599: 1, 3600: 1, 3601: 1, 3602: 1, 3603: 1, 3604: 10, 3605: 3,
			3606: 1, 3607: 1, 3608: 1, 3609: 9, 3610: 3, 3611: 4, 3612: 1, 3613: 2, 3614: 3, 3615: 2, 3616: 1, 3617: 3,
			3618: 5, 3619: 9, 3620: 2, 3621: 3, 3623: 3, 3624: 1, 3625: 2, 3626: 4, 3627: 4, 3628: 1, 3629: 6, 3630: 2,
			3631: 1, 3632: 3, 3633: 12, 3634: 19, 3635: 1, 3636: 6, 3637: 2, 3638: 1, 3639: 2, 3640: 3, 3641: 2, 3648: 7,
			3649: 1, 3650: 2, 3651: 3, 3653: 1, 3654: 2, 3655: 1, 3656: 10, 3657: 2, 3658: 1, 3659: 1, 3660: 2, 3663: 1}},
	{name: "Russian",
		input: "Съешь же ещё этих мягких французских булок да выпей чаю",
		want: map[rune]int{32: 9, 1057: 1, 1072: 3, 1073: 1, 1074: 1, 1075: 1, 1076: 1, 1077: 4, 1078: 1, 1079: 1, 1080: 3,
			1081: 1, 1082: 3, 1083: 1, 1084: 1, 1085: 1, 1086: 1, 1087: 1, 1088: 1, 1089: 1, 1090: 1, 1091: 2, 1092: 1, 1093: 3,
			1094: 1, 1095: 1, 1096: 1, 1097: 1, 1098: 1, 1099: 1, 1100: 1, 1101: 1, 1102: 1, 1103: 1, 1105: 1}},
	{name: "Emojis",
		input: "👾 🙇 💁 🙅 🙆 🙋 🙎 🐵 🙈 🙉 🙊 ❤️" +
			" 💔 💌 💕 💞 💓 💗 💖 💘 💝 💟 💜 💛 💚 💙",
		want: map[rune]int{32: 25, 10084: 1, 65039: 1, 128053: 1, 128126: 1, 128129: 1, 128140: 1,
			128147: 1, 128148: 1, 128149: 1, 128150: 1, 128151: 1, 128152: 1, 128153: 1, 128154: 1,
			128155: 1, 128156: 1, 128157: 1, 128158: 1, 128159: 1, 128581: 1, 128582: 1, 128583: 1,
			128584: 1, 128585: 1, 128586: 1, 128587: 1, 128590: 1}},
}

func TestLetters(t *testing.T) {
	for _, tt := range lettersTestData {
		tt := tt // as per sean- suggestion on this discussion https://github.com/kyoh86/scopelint/issues/4
		t.Run(tt.name, func(t *testing.T) {
			got := letters(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortLetters() = \n%v,\nwant \n%v", got, tt.want)
			}
		})
	}
}
