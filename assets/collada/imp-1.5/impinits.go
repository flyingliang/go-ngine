package collimp

import (
	xmlx "github.com/jteeuwen/go-pkg-xmlx"

	nga "github.com/go3d/go-ngine/assets"
)

func init_KxModelBinding(xn *xmlx.Node) (obj *nga.KxModelBinding) {
	obj = new(nga.KxModelBinding)

	load_KxModelBinding(xn, obj)
	return
}

func init_GeometryBrepCurves(xn *xmlx.Node) (obj *nga.GeometryBrepCurves) {
	obj = new(nga.GeometryBrepCurves)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepCurves(xn, obj)
	return
}

func init_Formula(xn *xmlx.Node) (obj *nga.Formula) {
	obj = new(nga.Formula)

	load_Formula(xn, obj)
	return
}

func init_Sources(xn *xmlx.Node) (obj *nga.Sources) {
	obj = new(nga.Sources)

	load_Sources(xn, obj)
	return
}

func init_PxSceneInst(xn *xmlx.Node) (obj *nga.PxSceneInst) {
	obj = new(nga.PxSceneInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_PxSceneInst(xn, obj)
	obj.SetDirty()
	return
}

func init_SidVec3(xn *xmlx.Node) (obj *nga.SidVec3) {
	obj = new(nga.SidVec3)
	has_Sid(xn, &obj.HasSid)

	load_SidVec3(xn, obj)
	return
}

func init_GeometryPrimitives(xn *xmlx.Node) (obj *nga.GeometryPrimitives) {
	obj = new(nga.GeometryPrimitives)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_GeometryPrimitives(xn, obj)
	return
}

func init_KxJointDef(xn *xmlx.Node) (obj *nga.KxJointDef) {
	obj = new(nga.KxJointDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_KxJointDef(xn, obj)
	obj.SetDirty()
	return
}

func init_KxFrameTcp(xn *xmlx.Node) (obj *nga.KxFrameTcp) {
	obj = new(nga.KxFrameTcp)

	load_KxFrameTcp(xn, obj)
	return
}

func init_FxPassEvaluationClearColor(xn *xmlx.Node) (obj *nga.FxPassEvaluationClearColor) {
	obj = new(nga.FxPassEvaluationClearColor)

	load_FxPassEvaluationClearColor(xn, obj)
	return
}

func init_FxCreateInitFrom(xn *xmlx.Node) (obj *nga.FxCreateInitFrom) {
	obj = new(nga.FxCreateInitFrom)

	load_FxCreateInitFrom(xn, obj)
	return
}

func init_FormulaInst(xn *xmlx.Node) (obj *nga.FormulaInst) {
	obj = new(nga.FormulaInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_ParamInsts(xn, &obj.HasParamInsts)
	has_Sid(xn, &obj.HasSid)

	load_FormulaInst(xn, obj)
	obj.SetDirty()
	return
}

func init_FxEffectInstTechniqueHint(xn *xmlx.Node) (obj *nga.FxEffectInstTechniqueHint) {
	obj = new(nga.FxEffectInstTechniqueHint)

	load_FxEffectInstTechniqueHint(xn, obj)
	return
}

func init_KxJointKind(xn *xmlx.Node) (obj *nga.KxJointKind) {
	obj = new(nga.KxJointKind)

	load_KxJointKind(xn, obj)
	return
}

func init_LightAmbient(xn *xmlx.Node) (obj *nga.LightAmbient) {
	obj = new(nga.LightAmbient)

	load_LightAmbient(xn, obj)
	return
}

func init_GeometryBrepOrientation(xn *xmlx.Node) (obj *nga.GeometryBrepOrientation) {
	obj = new(nga.GeometryBrepOrientation)

	load_GeometryBrepOrientation(xn, obj)
	return
}

func init_ChildNode(xn *xmlx.Node) (obj *nga.ChildNode) {
	obj = new(nga.ChildNode)

	load_ChildNode(xn, obj)
	return
}

func init_Transform(xn *xmlx.Node) (obj *nga.Transform) {
	obj = new(nga.Transform)
	has_Sid(xn, &obj.HasSid)

	load_Transform(xn, obj)
	return
}

func init_CameraOptics(xn *xmlx.Node) (obj *nga.CameraOptics) {
	obj = new(nga.CameraOptics)
	has_Extras(xn, &obj.HasExtras)
	has_Techniques(xn, &obj.HasTechniques)

	load_CameraOptics(xn, obj)
	return
}

func init_FxProfileCommon(xn *xmlx.Node) (obj *nga.FxProfileCommon) {
	obj = new(nga.FxProfileCommon)

	load_FxProfileCommon(xn, obj)
	return
}

func init_NodeDef(xn *xmlx.Node) (obj *nga.NodeDef) {
	obj = new(nga.NodeDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_NodeDef(xn, obj)
	obj.SetDirty()
	return
}

func init_FxGlslTechniques(xn *xmlx.Node) (obj *nga.FxGlslTechniques) {
	obj = new(nga.FxGlslTechniques)

	load_FxGlslTechniques(xn, obj)
	return
}

func init_CameraOrthographic(xn *xmlx.Node) (obj *nga.CameraOrthographic) {
	obj = new(nga.CameraOrthographic)

	load_CameraOrthographic(xn, obj)
	return
}

func init_Scene(xn *xmlx.Node) (obj *nga.Scene) {
	obj = new(nga.Scene)
	has_Extras(xn, &obj.HasExtras)

	load_Scene(xn, obj)
	return
}

func init_FxPassProgramBindUniform(xn *xmlx.Node) (obj *nga.FxPassProgramBindUniform) {
	obj = new(nga.FxPassProgramBindUniform)

	load_FxPassProgramBindUniform(xn, obj)
	return
}

func init_GeometryBrepNurbs(xn *xmlx.Node) (obj *nga.GeometryBrepNurbs) {
	obj = nga.NewGeometryBrepNurbs()
	has_Extras(xn, &obj.HasExtras)
	has_Sources(xn, &obj.HasSources)

	load_GeometryBrepNurbs(xn, obj)
	return
}

func init_GeometryBrepEllipse(xn *xmlx.Node) (obj *nga.GeometryBrepEllipse) {
	obj = new(nga.GeometryBrepEllipse)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepEllipse(xn, obj)
	return
}

func init_GeometryInst(xn *xmlx.Node) (obj *nga.GeometryInst) {
	obj = new(nga.GeometryInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_GeometryInst(xn, obj)
	obj.SetDirty()
	return
}

func init_GeometryPositioning(xn *xmlx.Node) (obj *nga.GeometryPositioning) {
	obj = new(nga.GeometryPositioning)

	load_GeometryPositioning(xn, obj)
	return
}

func init_PxForceFieldInst(xn *xmlx.Node) (obj *nga.PxForceFieldInst) {
	obj = new(nga.PxForceFieldInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_PxForceFieldInst(xn, obj)
	obj.SetDirty()
	return
}

func init_GeometryBrepEdges(xn *xmlx.Node) (obj *nga.GeometryBrepEdges) {
	obj = new(nga.GeometryBrepEdges)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_GeometryBrepEdges(xn, obj)
	return
}

func init_Bool2(xn *xmlx.Node) (obj *nga.Bool2) {
	obj = new(nga.Bool2)

	load_Bool2(xn, obj)
	return
}

func init_FxCubeFace(xn *xmlx.Node) (obj *nga.FxCubeFace) {
	obj = new(nga.FxCubeFace)

	load_FxCubeFace(xn, obj)
	return
}

func init_Int2(xn *xmlx.Node) (obj *nga.Int2) {
	obj = new(nga.Int2)

	load_Int2(xn, obj)
	return
}

func init_Float2x4(xn *xmlx.Node) (obj *nga.Float2x4) {
	obj = new(nga.Float2x4)

	load_Float2x4(xn, obj)
	return
}

func init_KxLink(xn *xmlx.Node) (obj *nga.KxLink) {
	obj = new(nga.KxLink)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_KxLink(xn, obj)
	return
}

func init_FxPassProgramBindAttribute(xn *xmlx.Node) (obj *nga.FxPassProgramBindAttribute) {
	obj = new(nga.FxPassProgramBindAttribute)

	load_FxPassProgramBindAttribute(xn, obj)
	return
}

func init_KxModelInst(xn *xmlx.Node) (obj *nga.KxModelInst) {
	obj = new(nga.KxModelInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_ParamDefs(xn, &obj.HasParamDefs)
	has_ParamInsts(xn, &obj.HasParamInsts)
	has_Sid(xn, &obj.HasSid)

	load_KxModelInst(xn, obj)
	obj.SetDirty()
	return
}

func init_ParamDef(xn *xmlx.Node) (obj *nga.ParamDef) {
	obj = new(nga.ParamDef)
	has_Sid(xn, &obj.HasSid)

	load_ParamDef(xn, obj)
	return
}

func init_FxFormatRange(xn *xmlx.Node) (obj *nga.FxFormatRange) {
	obj = new(nga.FxFormatRange)

	load_FxFormatRange(xn, obj)
	return
}

func init_GeometryMesh(xn *xmlx.Node) (obj *nga.GeometryMesh) {
	obj = nga.NewGeometryMesh()
	has_Extras(xn, &obj.HasExtras)
	has_Sources(xn, &obj.HasSources)

	load_GeometryMesh(xn, obj)
	return
}

func init_CameraImager(xn *xmlx.Node) (obj *nga.CameraImager) {
	obj = new(nga.CameraImager)
	has_Extras(xn, &obj.HasExtras)
	has_Techniques(xn, &obj.HasTechniques)

	load_CameraImager(xn, obj)
	return
}

func init_PxRigidBodyDef(xn *xmlx.Node) (obj *nga.PxRigidBodyDef) {
	obj = new(nga.PxRigidBodyDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)
	has_Techniques(xn, &obj.HasTechniques)

	load_PxRigidBodyDef(xn, obj)
	obj.SetDirty()
	return
}

func init_FxCreateMips(xn *xmlx.Node) (obj *nga.FxCreateMips) {
	obj = new(nga.FxCreateMips)

	load_FxCreateMips(xn, obj)
	return
}

func init_GeometryBrepSolids(xn *xmlx.Node) (obj *nga.GeometryBrepSolids) {
	obj = new(nga.GeometryBrepSolids)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_GeometryBrepSolids(xn, obj)
	return
}

func init_LightDirectional(xn *xmlx.Node) (obj *nga.LightDirectional) {
	obj = new(nga.LightDirectional)

	load_LightDirectional(xn, obj)
	return
}

func init_GeometryVertices(xn *xmlx.Node) (obj *nga.GeometryVertices) {
	obj = new(nga.GeometryVertices)
	has_Extras(xn, &obj.HasExtras)
	has_Inputs(xn, &obj.HasInputs)
	has_Name(xn, &obj.HasName)

	load_GeometryVertices(xn, obj)
	return
}

func init_Float2x3(xn *xmlx.Node) (obj *nga.Float2x3) {
	obj = new(nga.Float2x3)

	load_Float2x3(xn, obj)
	return
}

func init_GeometryBrepCircle(xn *xmlx.Node) (obj *nga.GeometryBrepCircle) {
	obj = new(nga.GeometryBrepCircle)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepCircle(xn, obj)
	return
}

func init_KxAxisIndex(xn *xmlx.Node) (obj *nga.KxAxisIndex) {
	obj = new(nga.KxAxisIndex)

	load_KxAxisIndex(xn, obj)
	return
}

func init_FxSamplerKind(xn *xmlx.Node) (obj *nga.FxSamplerKind) {
	obj = new(nga.FxSamplerKind)

	load_FxSamplerKind(xn, obj)
	return
}

func init_KxAttachment(xn *xmlx.Node) (obj *nga.KxAttachment) {
	obj = new(nga.KxAttachment)

	load_KxAttachment(xn, obj)
	return
}

func init_GeometryBrepSweptSurface(xn *xmlx.Node) (obj *nga.GeometryBrepSweptSurface) {
	obj = new(nga.GeometryBrepSweptSurface)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepSweptSurface(xn, obj)
	return
}

func init_ParamOrRefSid(xn *xmlx.Node) (obj *nga.ParamOrRefSid) {
	obj = new(nga.ParamOrRefSid)

	load_ParamOrRefSid(xn, obj)
	return
}

func init_ControllerDef(xn *xmlx.Node) (obj *nga.ControllerDef) {
	obj = new(nga.ControllerDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_ControllerDef(xn, obj)
	obj.SetDirty()
	return
}

func init_FxBinding(xn *xmlx.Node) (obj *nga.FxBinding) {
	obj = new(nga.FxBinding)

	load_FxBinding(xn, obj)
	return
}

func init_FxSamplerFiltering(xn *xmlx.Node) (obj *nga.FxSamplerFiltering) {
	obj = new(nga.FxSamplerFiltering)

	load_FxSamplerFiltering(xn, obj)
	return
}

func init_VisualSceneEvaluation(xn *xmlx.Node) (obj *nga.VisualSceneEvaluation) {
	obj = new(nga.VisualSceneEvaluation)
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_VisualSceneEvaluation(xn, obj)
	obj.SetDirty()
	return
}

func init_FxSamplerWrapping(xn *xmlx.Node) (obj *nga.FxSamplerWrapping) {
	obj = new(nga.FxSamplerWrapping)

	load_FxSamplerWrapping(xn, obj)
	return
}

func init_GeometryBrepSphere(xn *xmlx.Node) (obj *nga.GeometryBrepSphere) {
	obj = new(nga.GeometryBrepSphere)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepSphere(xn, obj)
	return
}

func init_SourceAccessor(xn *xmlx.Node) (obj *nga.SourceAccessor) {
	obj = nga.NewSourceAccessor()

	load_SourceAccessor(xn, obj)
	return
}

func init_FxCreateFormat(xn *xmlx.Node) (obj *nga.FxCreateFormat) {
	obj = new(nga.FxCreateFormat)

	load_FxCreateFormat(xn, obj)
	return
}

func init_FxColorOrTexture(xn *xmlx.Node) (obj *nga.FxColorOrTexture) {
	obj = new(nga.FxColorOrTexture)

	load_FxColorOrTexture(xn, obj)
	return
}

func init_FxPassState(xn *xmlx.Node) (obj *nga.FxPassState) {
	obj = new(nga.FxPassState)

	load_FxPassState(xn, obj)
	return
}

func init_VisualSceneInst(xn *xmlx.Node) (obj *nga.VisualSceneInst) {
	obj = new(nga.VisualSceneInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_VisualSceneInst(xn, obj)
	obj.SetDirty()
	return
}

func init_FxEffectDef(xn *xmlx.Node) (obj *nga.FxEffectDef) {
	obj = new(nga.FxEffectDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_FxParamDefs(xn, &obj.HasFxParamDefs)
	has_Name(xn, &obj.HasName)

	load_FxEffectDef(xn, obj)
	obj.SetDirty()
	return
}

func init_Source(xn *xmlx.Node) (obj *nga.Source) {
	obj = new(nga.Source)
	has_Asset(xn, &obj.HasAsset)
	has_Name(xn, &obj.HasName)
	has_Techniques(xn, &obj.HasTechniques)

	load_Source(xn, obj)
	return
}

func init_Param(xn *xmlx.Node) (obj *nga.Param) {
	obj = new(nga.Param)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_Param(xn, obj)
	return
}

func init_FormulaDef(xn *xmlx.Node) (obj *nga.FormulaDef) {
	obj = new(nga.FormulaDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_ParamDefs(xn, &obj.HasParamDefs)
	has_Sid(xn, &obj.HasSid)
	has_Techniques(xn, &obj.HasTechniques)

	load_FormulaDef(xn, obj)
	obj.SetDirty()
	return
}

func init_GeometrySpline(xn *xmlx.Node) (obj *nga.GeometrySpline) {
	obj = nga.NewGeometrySpline()
	has_Extras(xn, &obj.HasExtras)
	has_Sources(xn, &obj.HasSources)

	load_GeometrySpline(xn, obj)
	return
}

func init_GeometryBrepCylinder(xn *xmlx.Node) (obj *nga.GeometryBrepCylinder) {
	obj = new(nga.GeometryBrepCylinder)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepCylinder(xn, obj)
	return
}

func init_Float4x4(xn *xmlx.Node) (obj *nga.Float4x4) {
	obj = new(nga.Float4x4)

	load_Float4x4(xn, obj)
	return
}

func init_KxModelDef(xn *xmlx.Node) (obj *nga.KxModelDef) {
	obj = new(nga.KxModelDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Techniques(xn, &obj.HasTechniques)

	load_KxModelDef(xn, obj)
	obj.SetDirty()
	return
}

func init_VisualSceneRenderingMaterialInst(xn *xmlx.Node) (obj *nga.VisualSceneRenderingMaterialInst) {
	obj = new(nga.VisualSceneRenderingMaterialInst)
	has_Extras(xn, &obj.HasExtras)

	load_VisualSceneRenderingMaterialInst(xn, obj)
	return
}

func init_FxPassProgram(xn *xmlx.Node) (obj *nga.FxPassProgram) {
	obj = new(nga.FxPassProgram)

	load_FxPassProgram(xn, obj)
	return
}

func init_PxShape(xn *xmlx.Node) (obj *nga.PxShape) {
	obj = new(nga.PxShape)
	has_Extras(xn, &obj.HasExtras)

	load_PxShape(xn, obj)
	return
}

func init_PxCylinder(xn *xmlx.Node) (obj *nga.PxCylinder) {
	obj = new(nga.PxCylinder)
	has_Extras(xn, &obj.HasExtras)

	load_PxCylinder(xn, obj)
	return
}

func init_KxBinding(xn *xmlx.Node) (obj *nga.KxBinding) {
	obj = new(nga.KxBinding)

	load_KxBinding(xn, obj)
	return
}

func init_ParamOrFloat(xn *xmlx.Node) (obj *nga.ParamOrFloat) {
	obj = new(nga.ParamOrFloat)

	load_ParamOrFloat(xn, obj)
	return
}

func init_GeometryPrimitiveKind(xn *xmlx.Node) (obj *nga.GeometryPrimitiveKind) {
	obj = new(nga.GeometryPrimitiveKind)

	load_GeometryPrimitiveKind(xn, obj)
	return
}

func init_CameraPerspective(xn *xmlx.Node) (obj *nga.CameraPerspective) {
	obj = new(nga.CameraPerspective)

	load_CameraPerspective(xn, obj)
	return
}

func init_GeometryDef(xn *xmlx.Node) (obj *nga.GeometryDef) {
	obj = new(nga.GeometryDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_GeometryDef(xn, obj)
	obj.SetDirty()
	return
}

func init_Extra(xn *xmlx.Node) (obj *nga.Extra) {
	obj = new(nga.Extra)
	has_Asset(xn, &obj.HasAsset)
	has_Name(xn, &obj.HasName)
	has_Techniques(xn, &obj.HasTechniques)

	load_Extra(xn, obj)
	return
}

func init_InputShared(xn *xmlx.Node) (obj *nga.InputShared) {
	obj = new(nga.InputShared)

	load_InputShared(xn, obj)
	return
}

func init_FxMaterialInst(xn *xmlx.Node) (obj *nga.FxMaterialInst) {
	obj = new(nga.FxMaterialInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_FxMaterialInst(xn, obj)
	obj.SetDirty()
	return
}

func init_KxSceneInst(xn *xmlx.Node) (obj *nga.KxSceneInst) {
	obj = new(nga.KxSceneInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_ParamDefs(xn, &obj.HasParamDefs)
	has_ParamInsts(xn, &obj.HasParamInsts)
	has_Sid(xn, &obj.HasSid)

	load_KxSceneInst(xn, obj)
	obj.SetDirty()
	return
}

func init_FxTechniqueGlsl(xn *xmlx.Node) (obj *nga.FxTechniqueGlsl) {
	obj = new(nga.FxTechniqueGlsl)
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Sid(xn, &obj.HasSid)

	load_FxTechniqueGlsl(xn, obj)
	return
}

func init_AnimationSampler(xn *xmlx.Node) (obj *nga.AnimationSampler) {
	obj = new(nga.AnimationSampler)
	has_Inputs(xn, &obj.HasInputs)

	load_AnimationSampler(xn, obj)
	return
}

func init_FxMaterialDef(xn *xmlx.Node) (obj *nga.FxMaterialDef) {
	obj = new(nga.FxMaterialDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_FxMaterialDef(xn, obj)
	obj.SetDirty()
	return
}

func init_FxColor(xn *xmlx.Node) (obj *nga.FxColor) {
	obj = new(nga.FxColor)
	has_Sid(xn, &obj.HasSid)

	load_FxColor(xn, obj)
	return
}

func init_LightInst(xn *xmlx.Node) (obj *nga.LightInst) {
	obj = new(nga.LightInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_LightInst(xn, obj)
	obj.SetDirty()
	return
}

func init_GeometryBrepLine(xn *xmlx.Node) (obj *nga.GeometryBrepLine) {
	obj = new(nga.GeometryBrepLine)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepLine(xn, obj)
	return
}

func init_FxTechniqueCommon(xn *xmlx.Node) (obj *nga.FxTechniqueCommon) {
	obj = new(nga.FxTechniqueCommon)
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Sid(xn, &obj.HasSid)

	load_FxTechniqueCommon(xn, obj)
	return
}

func init_GeometryBrepShells(xn *xmlx.Node) (obj *nga.GeometryBrepShells) {
	obj = new(nga.GeometryBrepShells)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_GeometryBrepShells(xn, obj)
	return
}

func init_PxSceneDef(xn *xmlx.Node) (obj *nga.PxSceneDef) {
	obj = new(nga.PxSceneDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Techniques(xn, &obj.HasTechniques)

	load_PxSceneDef(xn, obj)
	obj.SetDirty()
	return
}

func init_LightSpot(xn *xmlx.Node) (obj *nga.LightSpot) {
	obj = nga.NewLightSpot()

	load_LightSpot(xn, obj)
	return
}

func init_KxJointInst(xn *xmlx.Node) (obj *nga.KxJointInst) {
	obj = new(nga.KxJointInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_KxJointInst(xn, obj)
	obj.SetDirty()
	return
}

func init_ParamInst(xn *xmlx.Node) (obj *nga.ParamInst) {
	obj = new(nga.ParamInst)

	load_ParamInst(xn, obj)
	return
}

func init_CameraInst(xn *xmlx.Node) (obj *nga.CameraInst) {
	obj = new(nga.CameraInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_CameraInst(xn, obj)
	obj.SetDirty()
	return
}

func init_FxTechnique(xn *xmlx.Node) (obj *nga.FxTechnique) {
	obj = new(nga.FxTechnique)
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Sid(xn, &obj.HasSid)

	load_FxTechnique(xn, obj)
	return
}

func init_FxCreate(xn *xmlx.Node) (obj *nga.FxCreate) {
	obj = new(nga.FxCreate)

	load_FxCreate(xn, obj)
	return
}

func init_PxRigidConstraintAttachment(xn *xmlx.Node) (obj *nga.PxRigidConstraintAttachment) {
	obj = new(nga.PxRigidConstraintAttachment)
	has_Extras(xn, &obj.HasExtras)

	load_PxRigidConstraintAttachment(xn, obj)
	return
}

func init_FxPassEvaluationClearStencil(xn *xmlx.Node) (obj *nga.FxPassEvaluationClearStencil) {
	obj = new(nga.FxPassEvaluationClearStencil)

	load_FxPassEvaluationClearStencil(xn, obj)
	return
}

func init_LightDef(xn *xmlx.Node) (obj *nga.LightDef) {
	obj = new(nga.LightDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Techniques(xn, &obj.HasTechniques)

	load_LightDef(xn, obj)
	obj.SetDirty()
	return
}

func init_GeometryBrepBox(xn *xmlx.Node) (obj *nga.GeometryBrepBox) {
	obj = new(nga.GeometryBrepBox)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepBox(xn, obj)
	return
}

func init_ControllerInputs(xn *xmlx.Node) (obj *nga.ControllerInputs) {
	obj = new(nga.ControllerInputs)
	has_Extras(xn, &obj.HasExtras)
	has_Inputs(xn, &obj.HasInputs)

	load_ControllerInputs(xn, obj)
	return
}

func init_FxPassEvaluationClearDepth(xn *xmlx.Node) (obj *nga.FxPassEvaluationClearDepth) {
	obj = new(nga.FxPassEvaluationClearDepth)

	load_FxPassEvaluationClearDepth(xn, obj)
	return
}

func init_GeometryBrepSurfaceCurves(xn *xmlx.Node) (obj *nga.GeometryBrepSurfaceCurves) {
	obj = new(nga.GeometryBrepSurfaceCurves)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepSurfaceCurves(xn, obj)
	return
}

func init_PxRigidBodyInst(xn *xmlx.Node) (obj *nga.PxRigidBodyInst) {
	obj = new(nga.PxRigidBodyInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)
	has_Techniques(xn, &obj.HasTechniques)

	load_PxRigidBodyInst(xn, obj)
	obj.SetDirty()
	return
}

func init_GeometryBrepNurbsSurface(xn *xmlx.Node) (obj *nga.GeometryBrepNurbsSurface) {
	obj = nga.NewGeometryBrepNurbsSurface()
	has_Extras(xn, &obj.HasExtras)
	has_Sources(xn, &obj.HasSources)

	load_GeometryBrepNurbsSurface(xn, obj)
	return
}

func init_KxFrame(xn *xmlx.Node) (obj *nga.KxFrame) {
	obj = new(nga.KxFrame)

	load_KxFrame(xn, obj)
	return
}

func init_FxImageInst(xn *xmlx.Node) (obj *nga.FxImageInst) {
	obj = new(nga.FxImageInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_FxImageInst(xn, obj)
	obj.SetDirty()
	return
}

func init_FxTextureOpaque(xn *xmlx.Node) (obj *nga.FxTextureOpaque) {
	obj = new(nga.FxTextureOpaque)

	load_FxTextureOpaque(xn, obj)
	return
}

func init_FxSamplerImage(xn *xmlx.Node) (obj *nga.FxSamplerImage) {
	obj = new(nga.FxSamplerImage)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_FxSamplerImage(xn, obj)
	return
}

func init_GeometryControlVertices(xn *xmlx.Node) (obj *nga.GeometryControlVertices) {
	obj = new(nga.GeometryControlVertices)
	has_Extras(xn, &obj.HasExtras)
	has_Inputs(xn, &obj.HasInputs)

	load_GeometryControlVertices(xn, obj)
	return
}

func init_Float2x2(xn *xmlx.Node) (obj *nga.Float2x2) {
	obj = new(nga.Float2x2)

	load_Float2x2(xn, obj)
	return
}

func init_FxCreate2DSizeRatio(xn *xmlx.Node) (obj *nga.FxCreate2DSizeRatio) {
	obj = new(nga.FxCreate2DSizeRatio)

	load_FxCreate2DSizeRatio(xn, obj)
	return
}

func init_VisualSceneDef(xn *xmlx.Node) (obj *nga.VisualSceneDef) {
	obj = new(nga.VisualSceneDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_VisualSceneDef(xn, obj)
	obj.SetDirty()
	return
}

func init_Int4x4(xn *xmlx.Node) (obj *nga.Int4x4) {
	obj = new(nga.Int4x4)

	load_Int4x4(xn, obj)
	return
}

func init_KxAttachmentKind(xn *xmlx.Node) (obj *nga.KxAttachmentKind) {
	obj = new(nga.KxAttachmentKind)

	load_KxAttachmentKind(xn, obj)
	return
}

func init_GeometryBrepFaces(xn *xmlx.Node) (obj *nga.GeometryBrepFaces) {
	obj = new(nga.GeometryBrepFaces)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_GeometryBrepFaces(xn, obj)
	return
}

func init_KxFrameObject(xn *xmlx.Node) (obj *nga.KxFrameObject) {
	obj = new(nga.KxFrameObject)

	load_KxFrameObject(xn, obj)
	return
}

func init_KxJointAxisBinding(xn *xmlx.Node) (obj *nga.KxJointAxisBinding) {
	obj = new(nga.KxJointAxisBinding)

	load_KxJointAxisBinding(xn, obj)
	return
}

func init_MaterialBinding(xn *xmlx.Node) (obj *nga.MaterialBinding) {
	obj = new(nga.MaterialBinding)
	has_Extras(xn, &obj.HasExtras)
	has_Techniques(xn, &obj.HasTechniques)

	load_MaterialBinding(xn, obj)
	return
}

func init_Float3(xn *xmlx.Node) (obj *nga.Float3) {
	obj = new(nga.Float3)

	load_Float3(xn, obj)
	return
}

func init_Int4(xn *xmlx.Node) (obj *nga.Int4) {
	obj = new(nga.Int4)

	load_Int4(xn, obj)
	return
}

func init_IndexedInputs(xn *xmlx.Node) (obj *nga.IndexedInputs) {
	obj = new(nga.IndexedInputs)

	load_IndexedInputs(xn, obj)
	return
}

func init_Int2x2(xn *xmlx.Node) (obj *nga.Int2x2) {
	obj = new(nga.Int2x2)

	load_Int2x2(xn, obj)
	return
}

func init_AssetContributor(xn *xmlx.Node) (obj *nga.AssetContributor) {
	obj = new(nga.AssetContributor)

	load_AssetContributor(xn, obj)
	return
}

func init_FxPassEvaluation(xn *xmlx.Node) (obj *nga.FxPassEvaluation) {
	obj = new(nga.FxPassEvaluation)

	load_FxPassEvaluation(xn, obj)
	return
}

func init_Bool4(xn *xmlx.Node) (obj *nga.Bool4) {
	obj = new(nga.Bool4)

	load_Bool4(xn, obj)
	return
}

func init_GeometryBrepCone(xn *xmlx.Node) (obj *nga.GeometryBrepCone) {
	obj = new(nga.GeometryBrepCone)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepCone(xn, obj)
	return
}

func init_Layers(xn *xmlx.Node) (obj *nga.Layers) {
	obj = new(nga.Layers)

	load_Layers(xn, obj)
	return
}

func init_Float3x2(xn *xmlx.Node) (obj *nga.Float3x2) {
	obj = new(nga.Float3x2)

	load_Float3x2(xn, obj)
	return
}

func init_PxRigidBodyDefs(xn *xmlx.Node) (obj *nga.PxRigidBodyDefs) {
	obj = new(nga.PxRigidBodyDefs)

	load_PxRigidBodyDefs(xn, obj)
	return
}

func init_PxRigidConstraintInst(xn *xmlx.Node) (obj *nga.PxRigidConstraintInst) {
	obj = new(nga.PxRigidConstraintInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_PxRigidConstraintInst(xn, obj)
	obj.SetDirty()
	return
}

func init_FxProfileGlslCodeInclude(xn *xmlx.Node) (obj *nga.FxProfileGlslCodeInclude) {
	obj = new(nga.FxProfileGlslCodeInclude)
	has_Sid(xn, &obj.HasSid)

	load_FxProfileGlslCodeInclude(xn, obj)
	return
}

func init_FxPassProgramShader(xn *xmlx.Node) (obj *nga.FxPassProgramShader) {
	obj = new(nga.FxPassProgramShader)
	has_Extras(xn, &obj.HasExtras)

	load_FxPassProgramShader(xn, obj)
	return
}

func init_ControllerMorph(xn *xmlx.Node) (obj *nga.ControllerMorph) {
	obj = nga.NewControllerMorph()
	has_Sources(xn, &obj.HasSources)

	load_ControllerMorph(xn, obj)
	return
}

func init_KxKinematicsSystem(xn *xmlx.Node) (obj *nga.KxKinematicsSystem) {
	obj = new(nga.KxKinematicsSystem)
	has_Techniques(xn, &obj.HasTechniques)

	load_KxKinematicsSystem(xn, obj)
	return
}

func init_PxRigidBodyCommon(xn *xmlx.Node) (obj *nga.PxRigidBodyCommon) {
	obj = new(nga.PxRigidBodyCommon)

	load_PxRigidBodyCommon(xn, obj)
	return
}

func init_PxMaterialDef(xn *xmlx.Node) (obj *nga.PxMaterialDef) {
	obj = new(nga.PxMaterialDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Techniques(xn, &obj.HasTechniques)

	load_PxMaterialDef(xn, obj)
	obj.SetDirty()
	return
}

func init_FxPass(xn *xmlx.Node) (obj *nga.FxPass) {
	obj = nga.NewFxPass()
	has_Extras(xn, &obj.HasExtras)
	has_Sid(xn, &obj.HasSid)

	load_FxPass(xn, obj)
	return
}

func init_Float4(xn *xmlx.Node) (obj *nga.Float4) {
	obj = new(nga.Float4)

	load_Float4(xn, obj)
	return
}

func init_KxJointLimits(xn *xmlx.Node) (obj *nga.KxJointLimits) {
	obj = new(nga.KxJointLimits)

	load_KxJointLimits(xn, obj)
	return
}

func init_AnimationDef(xn *xmlx.Node) (obj *nga.AnimationDef) {
	obj = new(nga.AnimationDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sources(xn, &obj.HasSources)

	load_AnimationDef(xn, obj)
	obj.SetDirty()
	return
}

func init_FxFilterKind(xn *xmlx.Node) (obj *nga.FxFilterKind) {
	obj = new(nga.FxFilterKind)

	load_FxFilterKind(xn, obj)
	return
}

func init_GeometryBrepHyperbola(xn *xmlx.Node) (obj *nga.GeometryBrepHyperbola) {
	obj = new(nga.GeometryBrepHyperbola)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepHyperbola(xn, obj)
	return
}

func init_ParamOrBool(xn *xmlx.Node) (obj *nga.ParamOrBool) {
	obj = new(nga.ParamOrBool)

	load_ParamOrBool(xn, obj)
	return
}

func init_FxSampler(xn *xmlx.Node) (obj *nga.FxSampler) {
	obj = nga.NewFxSampler()
	has_Extras(xn, &obj.HasExtras)

	load_FxSampler(xn, obj)
	return
}

func init_PxRigidConstraintDefs(xn *xmlx.Node) (obj *nga.PxRigidConstraintDefs) {
	obj = new(nga.PxRigidConstraintDefs)

	load_PxRigidConstraintDefs(xn, obj)
	return
}

func init_LightPoint(xn *xmlx.Node) (obj *nga.LightPoint) {
	obj = nga.NewLightPoint()

	load_LightPoint(xn, obj)
	return
}

func init_FxCreateFormatHint(xn *xmlx.Node) (obj *nga.FxCreateFormatHint) {
	obj = new(nga.FxCreateFormatHint)

	load_FxCreateFormatHint(xn, obj)
	return
}

func init_Asset(xn *xmlx.Node) (obj *nga.Asset) {
	obj = nga.NewAsset()
	has_Extras(xn, &obj.HasExtras)

	load_Asset(xn, obj)
	return
}

func init_KxSceneDef(xn *xmlx.Node) (obj *nga.KxSceneDef) {
	obj = new(nga.KxSceneDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_KxSceneDef(xn, obj)
	obj.SetDirty()
	return
}

func init_FxCreateCube(xn *xmlx.Node) (obj *nga.FxCreateCube) {
	obj = new(nga.FxCreateCube)

	load_FxCreateCube(xn, obj)
	return
}

func init_Float3x3(xn *xmlx.Node) (obj *nga.Float3x3) {
	obj = new(nga.Float3x3)

	load_Float3x3(xn, obj)
	return
}

func init_GeometryBrepCurve(xn *xmlx.Node) (obj *nga.GeometryBrepCurve) {
	obj = new(nga.GeometryBrepCurve)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_GeometryBrepCurve(xn, obj)
	return
}

func init_FxShaderStage(xn *xmlx.Node) (obj *nga.FxShaderStage) {
	obj = new(nga.FxShaderStage)

	load_FxShaderStage(xn, obj)
	return
}

func init_SidString(xn *xmlx.Node) (obj *nga.SidString) {
	obj = new(nga.SidString)
	has_Sid(xn, &obj.HasSid)

	load_SidString(xn, obj)
	return
}

func init_Bool3(xn *xmlx.Node) (obj *nga.Bool3) {
	obj = new(nga.Bool3)

	load_Bool3(xn, obj)
	return
}

func init_FxEffectInst(xn *xmlx.Node) (obj *nga.FxEffectInst) {
	obj = new(nga.FxEffectInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_ParamInsts(xn, &obj.HasParamInsts)
	has_Sid(xn, &obj.HasSid)

	load_FxEffectInst(xn, obj)
	obj.SetDirty()
	return
}

func init_VisualSceneRendering(xn *xmlx.Node) (obj *nga.VisualSceneRendering) {
	obj = nga.NewVisualSceneRendering()
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_VisualSceneRendering(xn, obj)
	return
}

func init_FxCreate3D(xn *xmlx.Node) (obj *nga.FxCreate3D) {
	obj = new(nga.FxCreate3D)

	load_FxCreate3D(xn, obj)
	return
}

func init_Float2(xn *xmlx.Node) (obj *nga.Float2) {
	obj = new(nga.Float2)

	load_Float2(xn, obj)
	return
}

func init_KxJoint(xn *xmlx.Node) (obj *nga.KxJoint) {
	obj = new(nga.KxJoint)
	has_Sid(xn, &obj.HasSid)

	load_KxJoint(xn, obj)
	return
}

func init_AnimationInst(xn *xmlx.Node) (obj *nga.AnimationInst) {
	obj = new(nga.AnimationInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_AnimationInst(xn, obj)
	obj.SetDirty()
	return
}

func init_GeometryBrep(xn *xmlx.Node) (obj *nga.GeometryBrep) {
	obj = nga.NewGeometryBrep()
	has_Extras(xn, &obj.HasExtras)
	has_Sources(xn, &obj.HasSources)

	load_GeometryBrep(xn, obj)
	return
}

func init_ParamOrUint(xn *xmlx.Node) (obj *nga.ParamOrUint) {
	obj = new(nga.ParamOrUint)

	load_ParamOrUint(xn, obj)
	return
}

func init_LightAttenuation(xn *xmlx.Node) (obj *nga.LightAttenuation) {
	obj = nga.NewLightAttenuation()

	load_LightAttenuation(xn, obj)
	return
}

func init_Float4x2(xn *xmlx.Node) (obj *nga.Float4x2) {
	obj = new(nga.Float4x2)

	load_Float4x2(xn, obj)
	return
}

func init_PxModelInst(xn *xmlx.Node) (obj *nga.PxModelInst) {
	obj = new(nga.PxModelInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_PxModelInst(xn, obj)
	obj.SetDirty()
	return
}

func init_NodeInst(xn *xmlx.Node) (obj *nga.NodeInst) {
	obj = new(nga.NodeInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_NodeInst(xn, obj)
	obj.SetDirty()
	return
}

func init_FxVertexInputBinding(xn *xmlx.Node) (obj *nga.FxVertexInputBinding) {
	obj = new(nga.FxVertexInputBinding)

	load_FxVertexInputBinding(xn, obj)
	return
}

func init_SidFloat(xn *xmlx.Node) (obj *nga.SidFloat) {
	obj = new(nga.SidFloat)
	has_Sid(xn, &obj.HasSid)

	load_SidFloat(xn, obj)
	return
}

func init_FxWrapKind(xn *xmlx.Node) (obj *nga.FxWrapKind) {
	obj = new(nga.FxWrapKind)

	load_FxWrapKind(xn, obj)
	return
}

func init_GeometryBrepSurfaces(xn *xmlx.Node) (obj *nga.GeometryBrepSurfaces) {
	obj = new(nga.GeometryBrepSurfaces)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepSurfaces(xn, obj)
	return
}

func init_FxImageDef(xn *xmlx.Node) (obj *nga.FxImageDef) {
	obj = new(nga.FxImageDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_FxImageDef(xn, obj)
	obj.SetDirty()
	return
}

func init_FxPassEvaluationTarget(xn *xmlx.Node) (obj *nga.FxPassEvaluationTarget) {
	obj = nga.NewFxPassEvaluationTarget()

	load_FxPassEvaluationTarget(xn, obj)
	return
}

func init_ControllerInst(xn *xmlx.Node) (obj *nga.ControllerInst) {
	obj = new(nga.ControllerInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_ControllerInst(xn, obj)
	obj.SetDirty()
	return
}

func init_Document(xn *xmlx.Node) (obj *nga.Document) {
	obj = new(nga.Document)
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)

	load_Document(xn, obj)
	return
}

func init_FxCreate2D(xn *xmlx.Node) (obj *nga.FxCreate2D) {
	obj = new(nga.FxCreate2D)

	load_FxCreate2D(xn, obj)
	return
}

func init_SidBool(xn *xmlx.Node) (obj *nga.SidBool) {
	obj = new(nga.SidBool)
	has_Sid(xn, &obj.HasSid)

	load_SidBool(xn, obj)
	return
}

func init_PxRigidConstraintLimit(xn *xmlx.Node) (obj *nga.PxRigidConstraintLimit) {
	obj = new(nga.PxRigidConstraintLimit)

	load_PxRigidConstraintLimit(xn, obj)
	return
}

func init_SidFloat3(xn *xmlx.Node) (obj *nga.SidFloat3) {
	obj = new(nga.SidFloat3)
	has_Sid(xn, &obj.HasSid)

	load_SidFloat3(xn, obj)
	return
}

func init_ParamDefs(xn *xmlx.Node) (obj *nga.ParamDefs) {
	obj = new(nga.ParamDefs)

	load_ParamDefs(xn, obj)
	return
}

func init_Int3(xn *xmlx.Node) (obj *nga.Int3) {
	obj = new(nga.Int3)

	load_Int3(xn, obj)
	return
}

func init_AnimSamplerBehavior(xn *xmlx.Node) (obj *nga.AnimSamplerBehavior) {
	obj = new(nga.AnimSamplerBehavior)

	load_AnimSamplerBehavior(xn, obj)
	return
}

func init_TransformKind(xn *xmlx.Node) (obj *nga.TransformKind) {
	obj = new(nga.TransformKind)

	load_TransformKind(xn, obj)
	return
}

func init_FxParamDefs(xn *xmlx.Node) (obj *nga.FxParamDefs) {
	obj = new(nga.FxParamDefs)

	load_FxParamDefs(xn, obj)
	return
}

func init_GeometryBrepSurface(xn *xmlx.Node) (obj *nga.GeometryBrepSurface) {
	obj = new(nga.GeometryBrepSurface)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_GeometryBrepSurface(xn, obj)
	return
}

func init_FxCreate2DSizeExact(xn *xmlx.Node) (obj *nga.FxCreate2DSizeExact) {
	obj = new(nga.FxCreate2DSizeExact)

	load_FxCreate2DSizeExact(xn, obj)
	return
}

func init_AssetGeographicLocation(xn *xmlx.Node) (obj *nga.AssetGeographicLocation) {
	obj = new(nga.AssetGeographicLocation)

	load_AssetGeographicLocation(xn, obj)
	return
}

func init_FxPassProgramShaderSources(xn *xmlx.Node) (obj *nga.FxPassProgramShaderSources) {
	obj = new(nga.FxPassProgramShaderSources)

	load_FxPassProgramShaderSources(xn, obj)
	return
}

func init_GeometryBrepTorus(xn *xmlx.Node) (obj *nga.GeometryBrepTorus) {
	obj = new(nga.GeometryBrepTorus)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepTorus(xn, obj)
	return
}

func init_PxModelDef(xn *xmlx.Node) (obj *nga.PxModelDef) {
	obj = new(nga.PxModelDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_PxModelDef(xn, obj)
	obj.SetDirty()
	return
}

func init_Int3x3(xn *xmlx.Node) (obj *nga.Int3x3) {
	obj = new(nga.Int3x3)

	load_Int3x3(xn, obj)
	return
}

func init_KxFrameOrigin(xn *xmlx.Node) (obj *nga.KxFrameOrigin) {
	obj = new(nga.KxFrameOrigin)

	load_KxFrameOrigin(xn, obj)
	return
}

func init_KxMotionSystem(xn *xmlx.Node) (obj *nga.KxMotionSystem) {
	obj = new(nga.KxMotionSystem)
	has_Techniques(xn, &obj.HasTechniques)

	load_KxMotionSystem(xn, obj)
	return
}

func init_KxKinematicsAxis(xn *xmlx.Node) (obj *nga.KxKinematicsAxis) {
	obj = nga.NewKxKinematicsAxis()
	has_Name(xn, &obj.HasName)
	has_ParamDefs(xn, &obj.HasParamDefs)
	has_Sid(xn, &obj.HasSid)

	load_KxKinematicsAxis(xn, obj)
	return
}

func init_SourceArray(xn *xmlx.Node) (obj *nga.SourceArray) {
	obj = new(nga.SourceArray)
	has_Name(xn, &obj.HasName)

	load_SourceArray(xn, obj)
	return
}

func init_ControllerSkin(xn *xmlx.Node) (obj *nga.ControllerSkin) {
	obj = nga.NewControllerSkin()
	has_Sources(xn, &obj.HasSources)

	load_ControllerSkin(xn, obj)
	return
}

func init_GeometryBrepPcurves(xn *xmlx.Node) (obj *nga.GeometryBrepPcurves) {
	obj = new(nga.GeometryBrepPcurves)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_GeometryBrepPcurves(xn, obj)
	return
}

func init_ParamInsts(xn *xmlx.Node) (obj *nga.ParamInsts) {
	obj = new(nga.ParamInsts)

	load_ParamInsts(xn, obj)
	return
}

func init_FxParamDef(xn *xmlx.Node) (obj *nga.FxParamDef) {
	obj = new(nga.FxParamDef)
	has_Sid(xn, &obj.HasSid)

	load_FxParamDef(xn, obj)
	return
}

func init_GeometryPolygonHole(xn *xmlx.Node) (obj *nga.GeometryPolygonHole) {
	obj = new(nga.GeometryPolygonHole)

	load_GeometryPolygonHole(xn, obj)
	return
}

func init_FxProfileGlsl(xn *xmlx.Node) (obj *nga.FxProfileGlsl) {
	obj = nga.NewFxProfileGlsl()

	load_FxProfileGlsl(xn, obj)
	return
}

func init_GeometryBrepParabola(xn *xmlx.Node) (obj *nga.GeometryBrepParabola) {
	obj = new(nga.GeometryBrepParabola)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepParabola(xn, obj)
	return
}

func init_KxMotionAxis(xn *xmlx.Node) (obj *nga.KxMotionAxis) {
	obj = nga.NewKxMotionAxis()
	has_Name(xn, &obj.HasName)
	has_ParamDefs(xn, &obj.HasParamDefs)
	has_ParamInsts(xn, &obj.HasParamInsts)
	has_Sid(xn, &obj.HasSid)

	load_KxMotionAxis(xn, obj)
	return
}

func init_CameraDef(xn *xmlx.Node) (obj *nga.CameraDef) {
	obj = new(nga.CameraDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_CameraDef(xn, obj)
	obj.SetDirty()
	return
}

func init_FxTechniqueKind(xn *xmlx.Node) (obj *nga.FxTechniqueKind) {
	obj = new(nga.FxTechniqueKind)

	load_FxTechniqueKind(xn, obj)
	return
}

func init_FxAnnotation(xn *xmlx.Node) (obj *nga.FxAnnotation) {
	obj = new(nga.FxAnnotation)
	has_Name(xn, &obj.HasName)

	load_FxAnnotation(xn, obj)
	return
}

func init_AnimationChannel(xn *xmlx.Node) (obj *nga.AnimationChannel) {
	obj = new(nga.AnimationChannel)

	load_AnimationChannel(xn, obj)
	return
}

func init_Float7(xn *xmlx.Node) (obj *nga.Float7) {
	obj = new(nga.Float7)

	load_Float7(xn, obj)
	return
}

func init_KxArticulatedSystemInst(xn *xmlx.Node) (obj *nga.KxArticulatedSystemInst) {
	obj = new(nga.KxArticulatedSystemInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_ParamDefs(xn, &obj.HasParamDefs)
	has_ParamInsts(xn, &obj.HasParamInsts)
	has_Sid(xn, &obj.HasSid)

	load_KxArticulatedSystemInst(xn, obj)
	obj.SetDirty()
	return
}

func init_ParamOrFloat2(xn *xmlx.Node) (obj *nga.ParamOrFloat2) {
	obj = new(nga.ParamOrFloat2)

	load_ParamOrFloat2(xn, obj)
	return
}

func init_FxProfile(xn *xmlx.Node) (obj *nga.FxProfile) {
	obj = new(nga.FxProfile)
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_FxParamDefs(xn, &obj.HasFxParamDefs)

	load_FxProfile(xn, obj)
	return
}

func init_FxCreate3DInitFrom(xn *xmlx.Node) (obj *nga.FxCreate3DInitFrom) {
	obj = new(nga.FxCreate3DInitFrom)

	load_FxCreate3DInitFrom(xn, obj)
	return
}

func init_PxMaterialInst(xn *xmlx.Node) (obj *nga.PxMaterialInst) {
	obj = new(nga.PxMaterialInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_PxMaterialInst(xn, obj)
	obj.SetDirty()
	return
}

func init_PxMaterial(xn *xmlx.Node) (obj *nga.PxMaterial) {
	obj = new(nga.PxMaterial)

	load_PxMaterial(xn, obj)
	return
}

func init_Float3x4(xn *xmlx.Node) (obj *nga.Float3x4) {
	obj = new(nga.Float3x4)

	load_Float3x4(xn, obj)
	return
}

func init_FxInitFrom(xn *xmlx.Node) (obj *nga.FxInitFrom) {
	obj = new(nga.FxInitFrom)

	load_FxInitFrom(xn, obj)
	return
}

func init_KxAxisLimits(xn *xmlx.Node) (obj *nga.KxAxisLimits) {
	obj = new(nga.KxAxisLimits)

	load_KxAxisLimits(xn, obj)
	return
}

func init_ParamOrSidFloat(xn *xmlx.Node) (obj *nga.ParamOrSidFloat) {
	obj = new(nga.ParamOrSidFloat)

	load_ParamOrSidFloat(xn, obj)
	return
}

func init_AnimationClipDef(xn *xmlx.Node) (obj *nga.AnimationClipDef) {
	obj = new(nga.AnimationClipDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_AnimationClipDef(xn, obj)
	obj.SetDirty()
	return
}

func init_KxEffector(xn *xmlx.Node) (obj *nga.KxEffector) {
	obj = nga.NewKxEffector()
	has_Name(xn, &obj.HasName)
	has_ParamDefs(xn, &obj.HasParamDefs)
	has_ParamInsts(xn, &obj.HasParamInsts)
	has_Sid(xn, &obj.HasSid)

	load_KxEffector(xn, obj)
	return
}

func init_FxFormatPrecision(xn *xmlx.Node) (obj *nga.FxFormatPrecision) {
	obj = new(nga.FxFormatPrecision)

	load_FxFormatPrecision(xn, obj)
	return
}

func init_AnimationClipInst(xn *xmlx.Node) (obj *nga.AnimationClipInst) {
	obj = new(nga.AnimationClipInst)
	obj.Init()
	setInstDefRef(xn, &obj.BaseInst)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)

	load_AnimationClipInst(xn, obj)
	obj.SetDirty()
	return
}

func init_Technique(xn *xmlx.Node) (obj *nga.Technique) {
	obj = new(nga.Technique)

	load_Technique(xn, obj)
	return
}

func init_KxArticulatedSystemDef(xn *xmlx.Node) (obj *nga.KxArticulatedSystemDef) {
	obj = new(nga.KxArticulatedSystemDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_KxArticulatedSystemDef(xn, obj)
	obj.SetDirty()
	return
}

func init_GeometryBrepPlane(xn *xmlx.Node) (obj *nga.GeometryBrepPlane) {
	obj = new(nga.GeometryBrepPlane)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepPlane(xn, obj)
	return
}

func init_ParamOrInt(xn *xmlx.Node) (obj *nga.ParamOrInt) {
	obj = new(nga.ParamOrInt)

	load_ParamOrInt(xn, obj)
	return
}

func init_FxFormatChannels(xn *xmlx.Node) (obj *nga.FxFormatChannels) {
	obj = new(nga.FxFormatChannels)

	load_FxFormatChannels(xn, obj)
	return
}

func init_GeometryBrepWires(xn *xmlx.Node) (obj *nga.GeometryBrepWires) {
	obj = new(nga.GeometryBrepWires)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)

	load_GeometryBrepWires(xn, obj)
	return
}

func init_PxRigidConstraintSpring(xn *xmlx.Node) (obj *nga.PxRigidConstraintSpring) {
	obj = nga.NewPxRigidConstraintSpring()

	load_PxRigidConstraintSpring(xn, obj)
	return
}

func init_FxSamplerStates(xn *xmlx.Node) (obj *nga.FxSamplerStates) {
	obj = nga.NewFxSamplerStates()
	has_Extras(xn, &obj.HasExtras)

	load_FxSamplerStates(xn, obj)
	return
}

func init_GeometryBrepCapsule(xn *xmlx.Node) (obj *nga.GeometryBrepCapsule) {
	obj = new(nga.GeometryBrepCapsule)
	has_Extras(xn, &obj.HasExtras)

	load_GeometryBrepCapsule(xn, obj)
	return
}

func init_KxFrameTip(xn *xmlx.Node) (obj *nga.KxFrameTip) {
	obj = new(nga.KxFrameTip)

	load_KxFrameTip(xn, obj)
	return
}

func init_FxImageInitFrom(xn *xmlx.Node) (obj *nga.FxImageInitFrom) {
	obj = new(nga.FxImageInitFrom)

	load_FxImageInitFrom(xn, obj)
	return
}

func init_PxForceFieldDef(xn *xmlx.Node) (obj *nga.PxForceFieldDef) {
	obj = new(nga.PxForceFieldDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Techniques(xn, &obj.HasTechniques)

	load_PxForceFieldDef(xn, obj)
	obj.SetDirty()
	return
}

func init_Float4x3(xn *xmlx.Node) (obj *nga.Float4x3) {
	obj = new(nga.Float4x3)

	load_Float4x3(xn, obj)
	return
}

func init_FxTexture(xn *xmlx.Node) (obj *nga.FxTexture) {
	obj = new(nga.FxTexture)
	has_Extras(xn, &obj.HasExtras)

	load_FxTexture(xn, obj)
	return
}

func init_Input(xn *xmlx.Node) (obj *nga.Input) {
	obj = new(nga.Input)

	load_Input(xn, obj)
	return
}

func init_PxRigidConstraintDef(xn *xmlx.Node) (obj *nga.PxRigidConstraintDef) {
	obj = new(nga.PxRigidConstraintDef)
	obj.Init()
	has_Asset(xn, &obj.HasAsset)
	has_Extras(xn, &obj.HasExtras)
	has_Name(xn, &obj.HasName)
	has_Sid(xn, &obj.HasSid)
	has_Techniques(xn, &obj.HasTechniques)

	load_PxRigidConstraintDef(xn, obj)
	obj.SetDirty()
	return
}

func init_FxCreateCubeInitFrom(xn *xmlx.Node) (obj *nga.FxCreateCubeInitFrom) {
	obj = new(nga.FxCreateCubeInitFrom)

	load_FxCreateCubeInitFrom(xn, obj)
	return
}
