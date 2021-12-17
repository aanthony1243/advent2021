import "csv"
import "strings"
import "experimental"

lil_txt = csv.from(csv: 
"
_value
000000000100
000000011110
000000010110
000000010111
000000010101
000000001111
000000000111
000000011100
000000010000
000000011001
000000000010
000000001010
",
mode: "raw")

big_txt = 
csv.from(file: "./3a.txt", mode: "raw")

// I decided that it's "cheating" to manipulate the raw data before it touches flux.  
// this means that data that isn't in the right format for CSV needs to be munged by flux prior to 
// pulling it into a table
// comment 1: this one is probably easier to the more binary-clever, of which I am not a member
// comment 2: this would be easier if I had more iteration available outside of table processing. 
// comment 3: file possible bug with int^int exponent operator
// comment 4: another possible bug in csv.from in cloud, it seems to look for the file and shouldn't even try.
// In OSS it correctly denies file opening.  I'm using vs-code and I can only connect it to an influxdb so I guess local 
// file access is not really possible with this tool  
// comment 5: what I wanted to do was get oxmatch inside of filterOXCO
// but the function would hang and timeout when I tried.  
// comment 6: the yields on line 1108 look fine but get me an internal error
func3b = (t=<-, ind="unk") => {
    r = t
    |> reduce(
        fn: (r, accumulator) => {
            bits = strings.split(v: r._value, t: "")
            return {
                total: 1 + accumulator.total, 
                b0: accumulator.b0 + int(v:bits[0]),
                b1: accumulator.b1 + int(v:bits[1]), 
                b2: accumulator.b2 + int(v:bits[2]),
                b3: accumulator.b3 + int(v:bits[3]),
                b4: accumulator.b4 + int(v:bits[4]),
                b5: accumulator.b5 + int(v:bits[5]),
                b6: accumulator.b6 + int(v:bits[6]),
                b7: accumulator.b7 + int(v:bits[7]),
                b8: accumulator.b8 + int(v:bits[8]),
                b9: accumulator.b9 + int(v:bits[9]),
                b10:accumulator.b10 + int(v:bits[10]), 
                b11:accumulator.b11 + int(v:bits[11]), 
            }
        },
        identity: {
        total: 0,
        b0: 0,
        b1: 0, 
        b2: 0, 
        b3: 0, 
        b4: 0, 
        b5: 0, 
        b6: 0, 
        b7: 0, 
        b8: 0, 
        b9: 0, 
        b10: 0, 
        b11: 0,  
        }
    )
    // happy that this works...using a constant here requires this function to run no more than once
    |> yield(name: ind)
    |> findRecord(
    fn: (key) => true,
    idx: 0
    )
    cut = int(v:r.total / 2)
    return [
        (if r.b0 >= r.total - r.b0  then "1" else "0"),
        (if r.b1 >= r.total - r.b1  then "1" else "0"),
        (if r.b2 >= r.total - r.b2  then "1" else "0"),
        (if r.b3 >= r.total - r.b3  then "1" else "0"),
        (if r.b4 >= r.total - r.b4  then "1" else "0"),
        (if r.b5 >= r.total - r.b5  then "1" else "0"),
        (if r.b6 >= r.total - r.b6  then "1" else "0"),
        (if r.b7 >= r.total - r.b7  then "1" else "0"),
        (if r.b8 >= r.total - r.b8  then "1" else "0"),
        (if r.b9 >= r.total - r.b9  then "1" else "0"),
        (if r.b10 >= r.total - r.b10 then "1" else "0"),
        (if r.b11 >= r.total - r.b11 then "1" else "0"),
    ]
}




filterOXCO = (table=<-, ind=0, inv=false, sel) => 
  table |> filter(fn: (r) => {
    digits = strings.split(v: r._value, t: "")
    return (sel[ind] == digits[ind] and not inv) or (sel[ind] != digits[ind] and inv)
  })


bin2dec = (table=<-) =>
table 
|> map(fn: (r) => {
  bits = strings.split(v: r._value, t: "")
  // woops it was little endian
  dec = 
    (if bits[0] == "1" then 2.0^11.0 else 0.0) + 
    (if bits[1] == "1"  then 2.0^10.0 else 0.0) + 
    (if bits[2] == "1"  then 2.0^9.0 else 0.0) + 
    (if bits[3] == "1"  then 2.0^8.0 else 0.0) + 
    (if bits[4] == "1"  then 2.0^7.0 else 0.0) + 
    (if bits[5] == "1"  then 2.0^6.0 else 0.0) + 
    (if bits[6] == "1"  then 2.0^5.0 else 0.0) + 
    (if bits[7] == "1"  then 2.0^4.0 else 0.0) + 
    (if bits[8] == "1"  then 2.0^3.0 else 0.0) + 
    (if bits[9] == "1"  then 2.0^2.0 else 0.0) + 
    (if bits[10] == "1" then 2.0^1.0 else 0.0) + 
    (if bits[11] == "1" then 2.0^0.0 else 0.0) 
  
    return {bin: r._value, dec: dec}

})

inv = false 
oxmatch0 = big_txt |> func3b(ind: "0") 
b0 = big_txt 
|> filterOXCO(ind:0,inv:inv,sel: oxmatch0)// |> yield(name: "b0")
oxmatch1 = b0 |> func3b(ind: "1")
b1 = b0 
|> filterOXCO(ind:1,inv:inv,sel: oxmatch1)// |> yield(name: "b1")
oxmatch2 = b1 |> func3b(ind: "2")
b2 = b1 
|> filterOXCO(ind:2,inv:inv,sel: oxmatch2) //|> yield(name: "b2")
oxmatch3 = b2 |> func3b(ind: "3")
b3 = b2 
|> filterOXCO(ind:3,inv:inv,sel: oxmatch3)// |> yield(name: "b3")
oxmatch4 = b3 |> func3b(ind: "4")
b4 = b3 
|> filterOXCO(ind:4,inv:inv,sel: oxmatch4) //|> yield(name: "b4")
oxmatch5 = b4 |> func3b(ind: "5")
b5 = b4 
|> filterOXCO(ind:5,inv:inv,sel: oxmatch5) //|> yield(name: "b5")
oxmatch6 = b5 |> func3b(ind: "6")
b6 = b5 
|> filterOXCO(ind:6,inv:inv,sel: oxmatch6)// |> yield(name: "b6")
oxmatch7 = b6 |> func3b(ind: "7")
b7 = b6  
|> filterOXCO(ind:7,inv:inv,sel: oxmatch7) //|> yield(name: "b7")
oxmatch8 = b7 |> func3b(ind: "8")
b8 = b7 
|> filterOXCO(ind:8,inv:inv,sel: oxmatch8) //|> yield(name: "b8")
oxmatch9 = b8 |> func3b(ind: "9")
b9 = b8 
|> filterOXCO(ind:9,inv:inv,sel: oxmatch9) //|> yield(name: "b9")
oxmatch10 = b9 |> func3b(ind: "10")
b10 = b9
|> filterOXCO(ind:10,inv:inv,sel: oxmatch10) //|> yield(name: "b10")
oxmatch11 = b10 |> func3b(ind: "11")
b11 = b10 
|> filterOXCO(ind:11,inv:inv,sel: oxmatch11) //|> yield(name: "b11")

//lil_txt |> yield(name:"lil")
// b0 |> yield(name: "b0")
// b1 |> yield(name: "b1")
// b2 |> yield(name: "b2")
// b3 |> yield(name: "b3")
// b4 |> yield(name: "b4")
// b5 |> yield(name: "b5")
// b6 |> yield(name: "b6")
// b7 |> yield(name: "b7")
// b8 |> yield(name: "b8")
// b9 |> yield(name: "b9")
// b10 |> yield(name: "b10")
b11 |> bin2dec() |> yield(name: "OXbin")

inv2 = true 
comatch0 = big_txt |> func3b(ind: "0a") 
bco0 = big_txt 
|> filterOXCO(ind:0,inv:inv2,sel: comatch0)// |> yield(name: "bco0")
comatch1 = bco0 |> func3b(ind: "1a")
bco1 = bco0 
|> filterOXCO(ind:1,inv:inv2,sel: comatch1) //|> yield(name: "bco1")
comatch2 = bco1 |> func3b(ind: "2a")
bco2 = bco1 
|> filterOXCO(ind:2,inv:inv2,sel: comatch2) //|> yield(name: "bco2")
comatch3 = bco2 |> func3b(ind: "3a")
bco3 = bco2 
|> filterOXCO(ind:3,inv:inv2,sel: comatch3)// |> yield(name: "bco3")
comatch4 = bco3 |> func3b(ind: "4a")
bco4 = bco3 
|> filterOXCO(ind:4,inv:inv2,sel: comatch4) //|> yield(name: "bco4")
comatch5 = bco4 |> func3b(ind: "5a")
bco5 = bco4 
|> filterOXCO(ind:5,inv:inv2,sel: comatch5) //|> yield(name: "bco5")
comatch6 = bco5 |> func3b(ind: "6a")
bco6 = bco5 
|> filterOXCO(ind:6,inv:inv2,sel: comatch6)// |> yield(name: "bco6")
comatch7 = bco6 |> func3b(ind: "7a")
bco7 = bco6  
|> filterOXCO(ind:7,inv:inv2,sel: comatch7) //|> yield(name: "bco7")
comatch8 = bco7 |> func3b(ind: "8a")
bco8 = bco7 
|> filterOXCO(ind:8,inv:inv2,sel: comatch8) //|> yield(name: "bco8")
comatch9 = bco8 |> func3b(ind: "9a")
bco9 = bco8 
|> filterOXCO(ind:9,inv:inv2,sel: comatch9)  //|> yield(name: "bco9")
// comatch10 = bco9 |> func3b(ind: "10a")
// bco10 = bco9
// |> filterOXCO(ind:10,inv:inv2,sel: comatch10) //|> yield(name: "bco10")
// comatch11 = bco10 |> func3b(ind: "11a")
// bco11 = bco10 
// |> filterOXCO(ind:11,inv:inv2,sel: comatch11) //|> yield(name: "bco11")

// lil_txt |> yield(name:"lil")
// // bco0 |> yield(name: "bco0")
// // b1 |> yield(name: "b1")
// // b2 |> yield(name: "b2")
// // b3 |> yield(name: "b3")
// // b4 |> yield(name: "b4")
// // b5 |> yield(name: "b5")
// // b6 |> yield(name: "b6")
// bco7 |> yield(name: "b7")
// bco8 |> yield(name: "b8")
// bco9 |> yield(name: "b9")
// bco10 |> yield(name: "b10")
// bco11 |> bin2dec() |> yield(name: "CObin")

bco8 |> bin2dec() |> yield(name: "CObin")

