package vm

// Code generated automatically; DO NOT EDIT

const (
	optrap                   bcop = 0
	opret                    bcop = 1
	opretk                   bcop = 2
	opretbk                  bcop = 3
	opretsk                  bcop = 4
	opretbhk                 bcop = 5
	opinit                   bcop = 6
	opbroadcast0k            bcop = 7
	opbroadcast1k            bcop = 8
	opfalse                  bcop = 9
	opnotk                   bcop = 10
	opandk                   bcop = 11
	opandnk                  bcop = 12
	opork                    bcop = 13
	opxork                   bcop = 14
	opxnork                  bcop = 15
	opcvtktof64              bcop = 16
	opcvtktoi64              bcop = 17
	opcvti64tok              bcop = 18
	opcvtf64tok              bcop = 19
	opcvti64tof64            bcop = 20
	opcvttruncf64toi64       bcop = 21
	opcvtfloorf64toi64       bcop = 22
	opcvtceilf64toi64        bcop = 23
	opcvti64tostr            bcop = 24
	opsortcmpvnf             bcop = 25
	opsortcmpvnl             bcop = 26
	opcmpv                   bcop = 27
	opcmpvk                  bcop = 28
	opcmpvkimm               bcop = 29
	opcmpvi64                bcop = 30
	opcmpvi64imm             bcop = 31
	opcmpvf64                bcop = 32
	opcmpvf64imm             bcop = 33
	opcmpltstr               bcop = 34
	opcmplestr               bcop = 35
	opcmpgtstr               bcop = 36
	opcmpgestr               bcop = 37
	opcmpltk                 bcop = 38
	opcmpltkimm              bcop = 39
	opcmplek                 bcop = 40
	opcmplekimm              bcop = 41
	opcmpgtk                 bcop = 42
	opcmpgtkimm              bcop = 43
	opcmpgek                 bcop = 44
	opcmpgekimm              bcop = 45
	opcmpeqf64               bcop = 46
	opcmpeqf64imm            bcop = 47
	opcmpltf64               bcop = 48
	opcmpltf64imm            bcop = 49
	opcmplef64               bcop = 50
	opcmplef64imm            bcop = 51
	opcmpgtf64               bcop = 52
	opcmpgtf64imm            bcop = 53
	opcmpgef64               bcop = 54
	opcmpgef64imm            bcop = 55
	opcmpeqi64               bcop = 56
	opcmpeqi64imm            bcop = 57
	opcmplti64               bcop = 58
	opcmplti64imm            bcop = 59
	opcmplei64               bcop = 60
	opcmplei64imm            bcop = 61
	opcmpgti64               bcop = 62
	opcmpgti64imm            bcop = 63
	opcmpgei64               bcop = 64
	opcmpgei64imm            bcop = 65
	opisnanf                 bcop = 66
	opchecktag               bcop = 67
	optypebits               bcop = 68
	opisnullv                bcop = 69
	opisnotnullv             bcop = 70
	opistruev                bcop = 71
	opisfalsev               bcop = 72
	opcmpeqslice             bcop = 73
	opcmpeqv                 bcop = 74
	opcmpeqvimm              bcop = 75
	opdateaddmonth           bcop = 76
	opdateaddmonthimm        bcop = 77
	opdateaddyear            bcop = 78
	opdateaddquarter         bcop = 79
	opdatediffmicrosecond    bcop = 80
	opdatediffparam          bcop = 81
	opdatediffmqy            bcop = 82
	opdateextractmicrosecond bcop = 83
	opdateextractmillisecond bcop = 84
	opdateextractsecond      bcop = 85
	opdateextractminute      bcop = 86
	opdateextracthour        bcop = 87
	opdateextractday         bcop = 88
	opdateextractdow         bcop = 89
	opdateextractdoy         bcop = 90
	opdateextractmonth       bcop = 91
	opdateextractquarter     bcop = 92
	opdateextractyear        bcop = 93
	opdatetounixepoch        bcop = 94
	opdatetounixmicro        bcop = 95
	opdatetruncmillisecond   bcop = 96
	opdatetruncsecond        bcop = 97
	opdatetruncminute        bcop = 98
	opdatetrunchour          bcop = 99
	opdatetruncday           bcop = 100
	opdatetruncdow           bcop = 101
	opdatetruncmonth         bcop = 102
	opdatetruncquarter       bcop = 103
	opdatetruncyear          bcop = 104
	opunboxts                bcop = 105
	opboxts                  bcop = 106
	opwidthbucketf64         bcop = 107
	opwidthbucketi64         bcop = 108
	optimebucketts           bcop = 109
	opgeohash                bcop = 110
	opgeohashimm             bcop = 111
	opgeotilex               bcop = 112
	opgeotiley               bcop = 113
	opgeotilees              bcop = 114
	opgeotileesimm           bcop = 115
	opgeodistance            bcop = 116
	opalloc                  bcop = 117
	opconcatstr              bcop = 118
	opfindsym                bcop = 119
	opfindsym2               bcop = 120
	opblendv                 bcop = 121
	opblendk                 bcop = 122
	opblendi64               bcop = 123
	opblendf64               bcop = 124
	opblendslice             bcop = 125
	opunpack                 bcop = 126
	opunsymbolize            bcop = 127
	opunboxktoi64            bcop = 128
	opunboxcoercef64         bcop = 129
	opunboxcoercei64         bcop = 130
	opunboxcvtf64            bcop = 131
	opunboxcvti64            bcop = 132
	opboxf64                 bcop = 133
	opboxi64                 bcop = 134
	opboxk                   bcop = 135
	opboxstr                 bcop = 136
	opboxlist                bcop = 137
	opmakelist               bcop = 138
	opmakestruct             bcop = 139
	ophashvalue              bcop = 140
	ophashvalueplus          bcop = 141
	ophashmember             bcop = 142
	ophashlookup             bcop = 143
	opaggandk                bcop = 144
	opaggork                 bcop = 145
	opaggsumf                bcop = 146
	opaggsumi                bcop = 147
	opaggminf                bcop = 148
	opaggmini                bcop = 149
	opaggmaxf                bcop = 150
	opaggmaxi                bcop = 151
	opaggandi                bcop = 152
	opaggori                 bcop = 153
	opaggxori                bcop = 154
	opaggcount               bcop = 155
	opaggbucket              bcop = 156
	opaggslotandk            bcop = 157
	opaggslotork             bcop = 158
	opaggslotsumf            bcop = 159
	opaggslotsumi            bcop = 160
	opaggslotavgf            bcop = 161
	opaggslotavgi            bcop = 162
	opaggslotminf            bcop = 163
	opaggslotmini            bcop = 164
	opaggslotmaxf            bcop = 165
	opaggslotmaxi            bcop = 166
	opaggslotandi            bcop = 167
	opaggslotori             bcop = 168
	opaggslotxori            bcop = 169
	opaggslotcount           bcop = 170
	//lint:ignore ST1003 opcode naming convention
	opaggslotcount_v2         bcop = 171
	oplitref                  bcop = 172
	opauxval                  bcop = 173
	opsplit                   bcop = 174
	optuple                   bcop = 175
	opmovk                    bcop = 176
	opzerov                   bcop = 177
	opmovv                    bcop = 178
	opmovvk                   bcop = 179
	opmovf64                  bcop = 180
	opmovi64                  bcop = 181
	opobjectsize              bcop = 182
	oparraysize               bcop = 183
	oparrayposition           bcop = 184
	opCmpStrEqCs              bcop = 185
	opCmpStrEqCi              bcop = 186
	opCmpStrEqUTF8Ci          bcop = 187
	opCmpStrFuzzyA3           bcop = 188
	opCmpStrFuzzyUnicodeA3    bcop = 189
	opHasSubstrFuzzyA3        bcop = 190
	opHasSubstrFuzzyUnicodeA3 bcop = 191
	opSkip1charLeft           bcop = 192
	opSkip1charRight          bcop = 193
	opSkipNcharLeft           bcop = 194
	opSkipNcharRight          bcop = 195
	opTrimWsLeft              bcop = 196
	opTrimWsRight             bcop = 197
	opTrim4charLeft           bcop = 198
	opTrim4charRight          bcop = 199
	opoctetlength             bcop = 200
	opcharlength              bcop = 201
	opSubstr                  bcop = 202
	opSplitPart               bcop = 203
	opContainsPrefixCs        bcop = 204
	opContainsPrefixCi        bcop = 205
	opContainsPrefixUTF8Ci    bcop = 206
	opContainsSuffixCs        bcop = 207
	opContainsSuffixCi        bcop = 208
	opContainsSuffixUTF8Ci    bcop = 209
	opContainsSubstrCs        bcop = 210
	opContainsSubstrCi        bcop = 211
	opContainsSubstrUTF8Ci    bcop = 212
	opEqPatternCs             bcop = 213
	opEqPatternCi             bcop = 214
	opEqPatternUTF8Ci         bcop = 215
	opContainsPatternCs       bcop = 216
	opContainsPatternCi       bcop = 217
	opContainsPatternUTF8Ci   bcop = 218
	opIsSubnetOfIP4           bcop = 219
	opDfaT6                   bcop = 220
	opDfaT7                   bcop = 221
	opDfaT8                   bcop = 222
	opDfaT6Z                  bcop = 223
	opDfaT7Z                  bcop = 224
	opDfaT8Z                  bcop = 225
	opDfaLZ                   bcop = 226
	opslower                  bcop = 227
	opsupper                  bcop = 228
	opaggapproxcount          bcop = 229
	opaggapproxcountmerge     bcop = 230
	opaggslotapproxcount      bcop = 231
	opaggslotapproxcountmerge bcop = 232
	oppowuintf64              bcop = 233
	opbroadcasti64            bcop = 234
	opabsi64                  bcop = 235
	opnegi64                  bcop = 236
	opsigni64                 bcop = 237
	opsquarei64               bcop = 238
	opbitnoti64               bcop = 239
	opbitcounti64             bcop = 240
	//lint:ignore ST1003 opcode naming convention
	opbitcounti64_v2 bcop = 241
	opaddi64         bcop = 242
	opaddi64imm      bcop = 243
	opsubi64         bcop = 244
	opsubi64imm      bcop = 245
	oprsubi64imm     bcop = 246
	opmuli64         bcop = 247
	opmuli64imm      bcop = 248
	opdivi64         bcop = 249
	opdivi64imm      bcop = 250
	oprdivi64imm     bcop = 251
	opmodi64         bcop = 252
	opmodi64imm      bcop = 253
	oprmodi64imm     bcop = 254
	opaddmuli64imm   bcop = 255
	opminvaluei64    bcop = 256
	opminvaluei64imm bcop = 257
	opmaxvaluei64    bcop = 258
	opmaxvaluei64imm bcop = 259
	opandi64         bcop = 260
	opandi64imm      bcop = 261
	opori64          bcop = 262
	opori64imm       bcop = 263
	opxori64         bcop = 264
	opxori64imm      bcop = 265
	opslli64         bcop = 266
	opslli64imm      bcop = 267
	opsrai64         bcop = 268
	opsrai64imm      bcop = 269
	opsrli64         bcop = 270
	opsrli64imm      bcop = 271
	opbroadcastf64   bcop = 272
	opabsf64         bcop = 273
	opnegf64         bcop = 274
	opsignf64        bcop = 275
	opsquaref64      bcop = 276
	oproundf64       bcop = 277
	oproundevenf64   bcop = 278
	optruncf64       bcop = 279
	opfloorf64       bcop = 280
	opceilf64        bcop = 281
	opaddf64         bcop = 282
	opaddf64imm      bcop = 283
	opsubf64         bcop = 284
	opsubf64imm      bcop = 285
	oprsubf64imm     bcop = 286
	opmulf64         bcop = 287
	opmulf64imm      bcop = 288
	opdivf64         bcop = 289
	opdivf64imm      bcop = 290
	oprdivf64imm     bcop = 291
	opmodf64         bcop = 292
	opmodf64imm      bcop = 293
	oprmodf64imm     bcop = 294
	opminvaluef64    bcop = 295
	opminvaluef64imm bcop = 296
	opmaxvaluef64    bcop = 297
	opmaxvaluef64imm bcop = 298
	opsqrtf64        bcop = 299
	opcbrtf64        bcop = 300
	opexpf64         bcop = 301
	opexp2f64        bcop = 302
	opexp10f64       bcop = 303
	opexpm1f64       bcop = 304
	oplnf64          bcop = 305
	opln1pf64        bcop = 306
	oplog2f64        bcop = 307
	oplog10f64       bcop = 308
	opsinf64         bcop = 309
	opcosf64         bcop = 310
	optanf64         bcop = 311
	opasinf64        bcop = 312
	opacosf64        bcop = 313
	opatanf64        bcop = 314
	opatan2f64       bcop = 315
	ophypotf64       bcop = 316
	oppowf64         bcop = 317
	_maxbcop              = 318
)

type opreplace struct{ from, to bcop }

var patchAVX512Level2 []opreplace = []opreplace{
	{from: opaggslotcount, to: opaggslotcount_v2},
	{from: opbitcounti64, to: opbitcounti64_v2},
}

// checksum: 18d404112e07dd75d04a3f358cebe587
