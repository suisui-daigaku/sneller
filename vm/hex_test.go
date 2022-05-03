// Copyright (C) 2022 Sneller, Inc.
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package vm_test

import (
	"bytes"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/SnellerInc/sneller/expr/partiql"
	"github.com/SnellerInc/sneller/plan"
)

func TestLiteral(t *testing.T) {
	tcs := []struct {
		query, body string
	}{
		{
			query: `select eventSource from input`,
			body: `e00100eaee05a88183de05a387be059f89617773526567696f6e87657665
6e744944896576656e7454696d65896576656e74547970658c6576656e74
56657273696f6e88686f73746e616d658e936b696e657369735061727469
74696f6e4b65798a72616e646f6d736565648e8f736f7572636549504164
647265737389757365724167656e748e8e696e736967687444657461696c
738361736e8567656f697089697061646472657373866e756d6265728e91
6f7267616e697a6174696f6e5f6e616d65846369747987636f756e747279
8c636f756e7472795f636f6465886c6f636174696f6e836c6174836c6f6e
896576656e744e616d658b6576656e74536f757263658c61747472696275
74696f6e738d6576656e7443617465676f72798e8e696e7369676874436f
6e746578748b696e73696768745479706585737461746588626173656c69
6e6587696e736967687487617665726167658576616c75658a7374617469
73746963738e90626173656c696e654475726174696f6e8e8f696e736967
68744475726174696f6e8e8f6d616e6167656d656e744576656e74887265
61644f6e6c7985666f7263658a696e7374616e63654964856974656d738c
696e7374616e6365735365748e9172657175657374506172616d65746572
7384636f64658c63757272656e7453746174658d70726576696f75735374
6174658e90726573706f6e7365456c656d656e7473896163636f756e7449
648b7072696e636970616c496484747970658c757365724964656e746974
7989696e766f6b656442798c6372656174696f6e446174658e906d666141
757468656e746963617465648e8e736f757263654964656e746974798a61
7474726962757465738e8e73657373696f6e436f6e74657874896572726f
72436f64658c6572726f724d65737361676588757365724e616d65de03b4
8a8d75732d676f762d656173742d318b8ea466363731373832662d363365
342d346265332d623762352d6532393065613362313062668c68800fe58a
9a89b7878d8e94415753436c6f7564547261696c496e73696768748e8431
2e30388f8c62386430366538326462636190813591483fed4fba4d6529a1
92de00e095de939822129e998d4b6f7265612054656c65636f6d96deb59a
8b4779656f6e6767692d646f9b8b536f757468204b6f7265619c824b529d
de949e484042d49ba5e353f89f48405fb5460aa64c30978e8f3137352e31
39342e3137312e323033938e01834d6f7a696c6c612f352e30202857696e
646f7773204e542031302e303b2057696e36343b2078363429204170706c
655765624b69742f3533372e333620284b48544d4c2c206c696b65204765
636b6f29204368726f6d652f39322e302e343531352e3135392053616661
72692f3533372e3336204564672f39322e302e3930322e373800c38ea647
656e657261746564457863657074696f6e415753436c6f7564547261696c
496e736967687400c48ea3416e206572726f72206f636375727265642064
7572696e672074686973206576656e74de03938a8c63612d63656e747261
6c2d318b8ea432366431356265322d303465332d343135652d623834362d
6531383437346431343161338c68800fe58a9a89b7878d8a417773417069
43616c6c8e84312e30388f8c62386430366538326462636190813791483f
c41d94d227ad5492de00d395dc9822904999864e656f74656c96deb59a8a
4d70756d616c616e67619b8c536f757468204166726963619c825a419dde
949e48c03976a7ef9db22d9f48403ef65fd8adab9f978b34312e3137302e
302e3636938e00d86177732d636c692f322e322e3520507974686f6e2f33
2e382e382057696e646f77732f3130206578652f414d4436342070726f6d
70742f6f666620636f6d6d616e642f69616d2e6372656174652d61636365
73732d6b6579a08e9053656e645353485075626c69634b6579a18ea26563
322d696e7374616e63652d636f6e6e6563742e616d617a6f6e6177732e63
6f6dae11af11bcdeb5b98c333435363738393132313537ba8e9541424344
454647484556535136433252414e443135bb8749414d5573657200c5836a
696d`,
		},
	}

	decode := func(str string) []byte {
		str = strings.Replace(str, "\n", "", -1)
		buf, err := hex.DecodeString(str)
		if err != nil {
			t.Fatal(err)
		}
		return buf
	}

	for i := range tcs {
		s, err := partiql.Parse([]byte(tcs[i].query))
		if err != nil {
			t.Fatal(err)
		}
		env := &queryenv{in: []plan.TableHandle{
			bufhandle(decode(tcs[i].body)),
		}}
		op, err := plan.New(s, env)
		if err != nil {
			t.Fatal(err)
		}
		var out bytes.Buffer
		var stats plan.ExecStats
		err = plan.Exec(op, &out, &stats)
		if err != nil {
			t.Fatal(err)
		}
	}
}
