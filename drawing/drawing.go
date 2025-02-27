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

package drawing ;import (_d "github.com/unidoc/unioffice";_b "github.com/unidoc/unioffice/color";_g "github.com/unidoc/unioffice/measurement";_bd "github.com/unidoc/unioffice/schema/soo/dml";);func (_ec ShapeProperties )ensureXfrm (){if _ec ._gfe .Xfrm ==nil {_ec ._gfe .Xfrm =_bd .NewCT_Transform2D ();};};func (_gd LineProperties )clearFill (){_gd ._e .NoFill =nil ;_gd ._e .GradFill =nil ;_gd ._e .SolidFill =nil ;_gd ._e .PattFill =nil ;};

// SetFlipHorizontal controls if the shape is flipped horizontally.
func (_adf ShapeProperties )SetFlipHorizontal (b bool ){_adf .ensureXfrm ();if !b {_adf ._gfe .Xfrm .FlipHAttr =nil ;}else {_adf ._gfe .Xfrm .FlipHAttr =_d .Bool (true );};};

// SetFlipVertical controls if the shape is flipped vertically.
func (_ega ShapeProperties )SetFlipVertical (b bool ){_ega .ensureXfrm ();if !b {_ega ._gfe .Xfrm .FlipVAttr =nil ;}else {_ega ._gfe .Xfrm .FlipVAttr =_d .Bool (true );};};

// MakeParagraphProperties constructs a new ParagraphProperties wrapper.
func MakeParagraphProperties (x *_bd .CT_TextParagraphProperties )ParagraphProperties {return ParagraphProperties {x };};

// MakeRunProperties constructs a new RunProperties wrapper.
func MakeRunProperties (x *_bd .CT_TextCharacterProperties )RunProperties {return RunProperties {x }};

// Properties returns the run's properties.
func (_gf Run )Properties ()RunProperties {if _gf ._fc .R ==nil {_gf ._fc .R =_bd .NewCT_RegularTextRun ();};if _gf ._fc .R .RPr ==nil {_gf ._fc .R .RPr =_bd .NewCT_TextCharacterProperties ();};return RunProperties {_gf ._fc .R .RPr };};

// SetHeight sets the height of the shape.
func (_cb ShapeProperties )SetHeight (h _g .Distance ){_cb .ensureXfrm ();if _cb ._gfe .Xfrm .Ext ==nil {_cb ._gfe .Xfrm .Ext =_bd .NewCT_PositiveSize2D ();};_cb ._gfe .Xfrm .Ext .CyAttr =int64 (h /_g .EMU );};

// GetPosition gets the position of the shape in EMU.
func (_aad ShapeProperties )GetPosition ()(int64 ,int64 ){_aad .ensureXfrm ();if _aad ._gfe .Xfrm .Off ==nil {_aad ._gfe .Xfrm .Off =_bd .NewCT_Point2D ();};return *_aad ._gfe .Xfrm .Off .XAttr .ST_CoordinateUnqualified ,*_aad ._gfe .Xfrm .Off .YAttr .ST_CoordinateUnqualified ;};

// MakeParagraph constructs a new paragraph wrapper.
func MakeParagraph (x *_bd .CT_TextParagraph )Paragraph {return Paragraph {x }};

// X returns the inner wrapped XML type.
func (_cf ParagraphProperties )X ()*_bd .CT_TextParagraphProperties {return _cf ._ca };

// X returns the inner wrapped XML type.
func (_ea Paragraph )X ()*_bd .CT_TextParagraph {return _ea ._ed };func (_a LineProperties )SetNoFill (){_a .clearFill ();_a ._e .NoFill =_bd .NewCT_NoFillProperties ()};

// AddBreak adds a new line break to a paragraph.
func (_ge Paragraph )AddBreak (){_geb :=_bd .NewEG_TextRun ();_geb .Br =_bd .NewCT_TextLineBreak ();_ge ._ed .EG_TextRun =append (_ge ._ed .EG_TextRun ,_geb );};func (_df ShapeProperties )SetNoFill (){_df .clearFill ();_df ._gfe .NoFill =_bd .NewCT_NoFillProperties ();};

// SetWidth sets the width of the shape.
func (_ac ShapeProperties )SetWidth (w _g .Distance ){_ac .ensureXfrm ();if _ac ._gfe .Xfrm .Ext ==nil {_ac ._gfe .Xfrm .Ext =_bd .NewCT_PositiveSize2D ();};_ac ._gfe .Xfrm .Ext .CxAttr =int64 (w /_g .EMU );};type LineProperties struct{_e *_bd .CT_LineProperties };func (_cc LineProperties )SetSolidFill (c _b .Color ){_cc .clearFill ();_cc ._e .SolidFill =_bd .NewCT_SolidColorFillProperties ();_cc ._e .SolidFill .SrgbClr =_bd .NewCT_SRgbColor ();_cc ._e .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// RunProperties controls the run properties.
type RunProperties struct{_ga *_bd .CT_TextCharacterProperties ;};

// SetAlign controls the paragraph alignment
func (_bec ParagraphProperties )SetAlign (a _bd .ST_TextAlignType ){_bec ._ca .AlgnAttr =a };

// Properties returns the paragraph properties.
func (_ff Paragraph )Properties ()ParagraphProperties {if _ff ._ed .PPr ==nil {_ff ._ed .PPr =_bd .NewCT_TextParagraphProperties ();};return MakeParagraphProperties (_ff ._ed .PPr );};

// SetFont controls the font of a run.
func (_gg RunProperties )SetFont (s string ){_gg ._ga .Latin =_bd .NewCT_TextFont ();_gg ._ga .Latin .TypefaceAttr =s ;};

// SetBulletChar sets the bullet character for the paragraph.
func (_ae ParagraphProperties )SetBulletChar (c string ){if c ==""{_ae ._ca .BuChar =nil ;}else {_ae ._ca .BuChar =_bd .NewCT_TextCharBullet ();_ae ._ca .BuChar .CharAttr =c ;};};

// X returns the inner wrapped XML type.
func (_be LineProperties )X ()*_bd .CT_LineProperties {return _be ._e };

// MakeRun constructs a new Run wrapper.
func MakeRun (x *_bd .EG_TextRun )Run {return Run {x }};

// SetSize sets the font size of the run text
func (_db RunProperties )SetSize (sz _g .Distance ){_db ._ga .SzAttr =_d .Int32 (int32 (sz /_g .HundredthPoint ));};

// SetBold controls the bolding of a run.
func (_af RunProperties )SetBold (b bool ){_af ._ga .BAttr =_d .Bool (b )};

// SetBulletFont controls the font for the bullet character.
func (_ab ParagraphProperties )SetBulletFont (f string ){if f ==""{_ab ._ca .BuFont =nil ;}else {_ab ._ca .BuFont =_bd .NewCT_TextFont ();_ab ._ca .BuFont .TypefaceAttr =f ;};};

// SetSolidFill controls the text color of a run.
func (_afc RunProperties )SetSolidFill (c _b .Color ){_afc ._ga .NoFill =nil ;_afc ._ga .BlipFill =nil ;_afc ._ga .GradFill =nil ;_afc ._ga .GrpFill =nil ;_afc ._ga .PattFill =nil ;_afc ._ga .SolidFill =_bd .NewCT_SolidColorFillProperties ();_afc ._ga .SolidFill .SrgbClr =_bd .NewCT_SRgbColor ();_afc ._ga .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};func (_aa ShapeProperties )clearFill (){_aa ._gfe .NoFill =nil ;_aa ._gfe .BlipFill =nil ;_aa ._gfe .GradFill =nil ;_aa ._gfe .GrpFill =nil ;_aa ._gfe .SolidFill =nil ;_aa ._gfe .PattFill =nil ;};

// SetGeometry sets the shape type of the shape
func (_aff ShapeProperties )SetGeometry (g _bd .ST_ShapeType ){if _aff ._gfe .PrstGeom ==nil {_aff ._gfe .PrstGeom =_bd .NewCT_PresetGeometry2D ();};_aff ._gfe .PrstGeom .PrstAttr =g ;};

// ParagraphProperties allows controlling paragraph properties.
type ParagraphProperties struct{_ca *_bd .CT_TextParagraphProperties ;};

// SetWidth sets the line width, MS products treat zero as the minimum width
// that can be displayed.
func (_c LineProperties )SetWidth (w _g .Distance ){_c ._e .WAttr =_d .Int32 (int32 (w /_g .EMU ))};

// SetText sets the run's text contents.
func (_ead Run )SetText (s string ){_ead ._fc .Br =nil ;_ead ._fc .Fld =nil ;if _ead ._fc .R ==nil {_ead ._fc .R =_bd .NewCT_RegularTextRun ();};_ead ._fc .R .T =s ;};

// SetPosition sets the position of the shape.
func (_dg ShapeProperties )SetPosition (x ,y _g .Distance ){_dg .ensureXfrm ();if _dg ._gfe .Xfrm .Off ==nil {_dg ._gfe .Xfrm .Off =_bd .NewCT_Point2D ();};_dg ._gfe .Xfrm .Off .XAttr .ST_CoordinateUnqualified =_d .Int64 (int64 (x /_g .EMU ));_dg ._gfe .Xfrm .Off .YAttr .ST_CoordinateUnqualified =_d .Int64 (int64 (y /_g .EMU ));};

// Run is a run within a paragraph.
type Run struct{_fc *_bd .EG_TextRun };

// X returns the inner wrapped XML type.
func (_aeb Run )X ()*_bd .EG_TextRun {return _aeb ._fc };

// SetSize sets the width and height of the shape.
func (_dea ShapeProperties )SetSize (w ,h _g .Distance ){_dea .SetWidth (w );_dea .SetHeight (h )};

// X returns the inner wrapped XML type.
func (_gaf ShapeProperties )X ()*_bd .CT_ShapeProperties {return _gaf ._gfe };const (LineJoinRound LineJoin =iota ;LineJoinBevel ;LineJoinMiter ;);

// SetJoin sets the line join style.
func (_cg LineProperties )SetJoin (e LineJoin ){_cg ._e .Round =nil ;_cg ._e .Miter =nil ;_cg ._e .Bevel =nil ;switch e {case LineJoinRound :_cg ._e .Round =_bd .NewCT_LineJoinRound ();case LineJoinBevel :_cg ._e .Bevel =_bd .NewCT_LineJoinBevel ();case LineJoinMiter :_cg ._e .Miter =_bd .NewCT_LineJoinMiterProperties ();};};

// SetLevel sets the level of indentation of a paragraph.
func (_ef ParagraphProperties )SetLevel (idx int32 ){_ef ._ca .LvlAttr =_d .Int32 (idx )};

// AddRun adds a new run to a paragraph.
func (_ad Paragraph )AddRun ()Run {_de :=MakeRun (_bd .NewEG_TextRun ());_ad ._ed .EG_TextRun =append (_ad ._ed .EG_TextRun ,_de .X ());return _de ;};func MakeShapeProperties (x *_bd .CT_ShapeProperties )ShapeProperties {return ShapeProperties {x }};func (_cfd ShapeProperties )SetSolidFill (c _b .Color ){_cfd .clearFill ();_cfd ._gfe .SolidFill =_bd .NewCT_SolidColorFillProperties ();_cfd ._gfe .SolidFill .SrgbClr =_bd .NewCT_SRgbColor ();_cfd ._gfe .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// SetNumbered controls if bullets are numbered or not.
func (_eg ParagraphProperties )SetNumbered (scheme _bd .ST_TextAutonumberScheme ){if scheme ==_bd .ST_TextAutonumberSchemeUnset {_eg ._ca .BuAutoNum =nil ;}else {_eg ._ca .BuAutoNum =_bd .NewCT_TextAutonumberBullet ();_eg ._ca .BuAutoNum .TypeAttr =scheme ;};};

// LineJoin is the type of line join
type LineJoin byte ;type ShapeProperties struct{_gfe *_bd .CT_ShapeProperties };

// Paragraph is a paragraph within a document.
type Paragraph struct{_ed *_bd .CT_TextParagraph };func (_aef ShapeProperties )LineProperties ()LineProperties {if _aef ._gfe .Ln ==nil {_aef ._gfe .Ln =_bd .NewCT_LineProperties ();};return LineProperties {_aef ._gfe .Ln };};