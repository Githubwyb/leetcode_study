package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		words      []string
		target     string
		startIndex int
		Want       int
	}

	testGroup := []testCase{
		{[]string{"lwgdugypkmsvxwhwbrccrbtemvudwhctnaaonednsbodptendi", "lumylagwxpmmivpujfawgvdbtxpluwekdpeakrqelpvrflnsnr", "lngqwiijizfzzhlvvszaownnachqunlktsnhgsjeluvcpmavuj", "nabbqiyarxmzleesxrfaynypfxonyhvwhebfjsxyinepggxnns", "oiqghjtvrhpgvsjlzmrwbwpmomqqliqytdzawhkwematskflgf", "dtiwjpdgcsmhaiwxrgligxlibfmvclobjjhljrqlvpuiufskix", "svqgvooghuqszkrmcrayqclotygdqnxfetdrfrfvbypgiizzdk", "qzrmfzdiodkdbhwifsinmamljlztwqtaoivufkcyeydsvpmdzw", "ihaekjyxvnmhvtanysybyqvrtmffpkgmnxisgmmhkhbtonlwgo", "ucrvwdlifpckbimcvevgsniepuewjqpakwnxksumgvrnmhmtuk", "lngqwiijizfzzhlvvszaownnachqunlktsnhgsjeluvcpmavuj", "lngqwiijizfzzhlvvszaownnachqunlktsnhgsjeluvcpmavuj", "vdtvcclyyraevotgikgojlbefpfmlzypychxehnglgettackoz", "qxspwpzxfxyxalrjuliwtbyllamkifbknnhzyfeabavwvvdkzk", "vdtvcclyyraevotgikgojlbefpfmlzypychxehnglgettackoz", "ucrvwdlifpckbimcvevgsniepuewjqpakwnxksumgvrnmhmtuk", "dtiwjpdgcsmhaiwxrgligxlibfmvclobjjhljrqlvpuiufskix", "vtbfahotrkxwcfwzlijfoqdkrvdmavpllbcfvibrqeuntsmrcs", "mfhqksxfeeltpiupaijavavbpphjxyuzqlqehirxnmviiaqnfv", "oowbnwbktlmsawtbilyvksqkbuohhjxqbepxgklkrwpjzcvipe", "mpnnvwuqbczvmrwhzvsmtuumwjczqehuumjvfbpgoxlurjbckq", "byaschxhjcgnnodazzpisqriyszffaqqiwljbuigdvzzobrlmc", "dmctcimgeztojrvqwyygmnzfwtsrmmbgednmytsludnrpjjjvk", "qxspwpzxfxyxalrjuliwtbyllamkifbknnhzyfeabavwvvdkzk", "cawzflwjjopbemxqazpsrsrlxljwqlvzpvjbjarqeqgythrsou", "ydqjqvalipkkrcbdprgjjangclswdjpaajiwhxeupidxirlith", "lwgdugypkmsvxwhwbrccrbtemvudwhctnaaonednsbodptendi", "ejtkmbyqtwrlhwynnqggpjaaaydjqnczhtyphfgugpbardzlvj", "cawzflwjjopbemxqazpsrsrlxljwqlvzpvjbjarqeqgythrsou", "oowbnwbktlmsawtbilyvksqkbuohhjxqbepxgklkrwpjzcvipe", "khhwlijglujhgimvfvuwgggigppichzscdtsiklalgqeqsencq", "khhwlijglujhgimvfvuwgggigppichzscdtsiklalgqeqsencq", "lngqwiijizfzzhlvvszaownnachqunlktsnhgsjeluvcpmavuj", "frdsoraagsllmgtatzybegxotrtgsuwfzpzxkpegggvpodnhrr", "ynuartuisymsqxxdqwndonpqhczqpuekwbayfidfisjpriqogx", "shmpixycafoqskoegqfvsrvboiegqwlbrkiuoeetncdxqlqsov", "oiqghjtvrhpgvsjlzmrwbwpmomqqliqytdzawhkwematskflgf", "xjtatoefvpwwgsvmepltadmzngxtnahqypfxgjppbqsmqnjvxm", "vtbfahotrkxwcfwzlijfoqdkrvdmavpllbcfvibrqeuntsmrcs", "dmctcimgeztojrvqwyygmnzfwtsrmmbgednmytsludnrpjjjvk", "dsohnrdxdqrbwdjhqpphwvlzfyizqyoedckthdlhmkxjxewubc", "qriythowwswwwucgwmezpqqneatemdxfqzpeytlozzojguywby", "lyhmqyjnztwzqotqkpmvpueyzjfroousgkkerqvmwjnjtmlkuf", "qzrmfzdiodkdbhwifsinmamljlztwqtaoivufkcyeydsvpmdzw", "qzrmfzdiodkdbhwifsinmamljlztwqtaoivufkcyeydsvpmdzw", "gxrtwoayoosijaddrlbvxqsvbbvaziqemcpxgljlnexvjnnakk", "mjftbskulmuztejkopyftjsdeoyuvhvqragbkqlfhgqqkafvau", "mjftbskulmuztejkopyftjsdeoyuvhvqragbkqlfhgqqkafvau", "qzrmfzdiodkdbhwifsinmamljlztwqtaoivufkcyeydsvpmdzw", "qxspwpzxfxyxalrjuliwtbyllamkifbknnhzyfeabavwvvdkzk"}, "ydqjqvalipkkrcbdprgjjangclswdjpaajiwhxeupidxirlith",
			0, 25},
		{[]string{"hello", "i", "am", "leetcode", "hello"}, "hello", 1, 1},
		{[]string{"a", "b", "leetcode"}, "leetcode", 0, 1},
		{[]string{"i", "eat", "leetcode"}, "ate", 0, -1},
	}

	for i, v := range testGroup {
		result := closetTarget(v.words, v.target, v.startIndex)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := closetTarget1(v.words, v.target, v.startIndex)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
