//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package wildcard ;func MatchSimple (pattern ,name string )bool {if pattern ==""{return name ==pattern ;};if pattern =="\u002a"{return true ;};_b :=make ([]rune ,0,len (name ));_a :=make ([]rune ,0,len (pattern ));for _ ,_d :=range name {_b =append (_b ,_d );};for _ ,_e :=range pattern {_a =append (_a ,_e );};_cf :=true ;return _ac (_b ,_a ,_cf );};func _ac (_fdg ,_cd []rune ,_fb bool )bool {for len (_cd )> 0{switch _cd [0]{default:if len (_fdg )==0||_fdg [0]!=_cd [0]{return false ;};case '?':if len (_fdg )==0&&!_fb {return false ;};case '*':return _ac (_fdg ,_cd [1:],_fb )||(len (_fdg )> 0&&_ac (_fdg [1:],_cd ,_fb ));};_fdg =_fdg [1:];_cd =_cd [1:];};return len (_fdg )==0&&len (_cd )==0;};func Index (pattern ,name string )(_ce int ){if pattern ==""||pattern =="\u002a"{return 0;};_eb :=make ([]rune ,0,len (name ));_af :=make ([]rune ,0,len (pattern ));for _ ,_ad :=range name {_eb =append (_eb ,_ad );};for _ ,_ca :=range pattern {_af =append (_af ,_ca );};return _aa (_eb ,_af ,0);};func Match (pattern ,name string )(_ec bool ){if pattern ==""{return name ==pattern ;};if pattern =="\u002a"{return true ;};_cg :=make ([]rune ,0,len (name ));_db :=make ([]rune ,0,len (pattern ));for _ ,_ag :=range name {_cg =append (_cg ,_ag );};for _ ,_g :=range pattern {_db =append (_db ,_g );};_dc :=false ;return _ac (_cg ,_db ,_dc );};func _aa (_fdd ,_ef []rune ,_bf int )int {for len (_ef )> 0{switch _ef [0]{default:if len (_fdd )==0{return -1;};if _fdd [0]!=_ef [0]{return _aa (_fdd [1:],_ef ,_bf +1);};case '?':if len (_fdd )==0{return -1;};case '*':if len (_fdd )==0{return -1;};_ee :=_aa (_fdd ,_ef [1:],_bf );if _ee !=-1{return _bf ;}else {_ee =_aa (_fdd [1:],_ef ,_bf );if _ee !=-1{return _bf ;}else {return -1;};};};_fdd =_fdd [1:];_ef =_ef [1:];};return _bf ;};