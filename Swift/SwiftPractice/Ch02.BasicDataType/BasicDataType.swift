// use playgroud maybe more convenient

// 127
var maxInt8 = Int8.max
// 128
var minInt8 = Int8.min
// 32767
var maxInt16 = Int16.max
var minInt16 = Int16.min
var maxInt32 = Int32.max
var minInt32 = Int32.min
var maxInt64 = Int64.max
var minInt64 = Int64.min

if maxInt8 < 10 {
    print()
}

var dic2 = Dictionary(dictionaryLiteral: (1, "1"), (2, "2"), (3, "3"))

for item in dic2.keys.sorted(by: >) {
    print(dic2[item]!)
}

var q: Int? = 8
var value = q ?? 0
// value = (q != nil) ? q! : 0
// if let tmp = q {
//     print("q: \(tmp)")
// } else {
//     print("q not found")
// }
