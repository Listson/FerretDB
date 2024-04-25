// Code generated by "stringer -linecomment -type ErrorCode"; DO NOT EDIT.

package commonerrors

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[errUnset-0]
	_ = x[errInternalError-1]
	_ = x[ErrBadValue-2]
	_ = x[ErrFailedToParse-9]
	_ = x[ErrTypeMismatch-14]
	_ = x[ErrNamespaceNotFound-26]
	_ = x[ErrIndexNotFound-27]
	_ = x[ErrUnsuitableValueType-28]
	_ = x[ErrConflictingUpdateOperators-40]
	_ = x[ErrCursorNotFound-43]
	_ = x[ErrNamespaceExists-48]
	_ = x[ErrDollarPrefixedFieldName-52]
	_ = x[ErrInvalidID-53]
	_ = x[ErrEmptyName-56]
	_ = x[ErrCommandNotFound-59]
	_ = x[ErrImmutableField-66]
	_ = x[ErrCannotCreateIndex-67]
	_ = x[ErrInvalidOptions-72]
	_ = x[ErrInvalidNamespace-73]
	_ = x[ErrIndexOptionsConflict-85]
	_ = x[ErrIndexKeySpecsConflict-86]
	_ = x[ErrOperationFailed-96]
	_ = x[ErrDocumentValidationFailure-121]
	_ = x[ErrNotImplemented-238]
	_ = x[ErrMechanismUnavailable-334]
	_ = x[ErrDuplicateKey-11000]
	_ = x[ErrStageGroupInvalidFields-15947]
	_ = x[ErrStageGroupID-15948]
	_ = x[ErrStageGroupMissingID-15955]
	_ = x[ErrStageLimitZero-15958]
	_ = x[ErrMatchBadExpression-15959]
	_ = x[ErrProjectBadExpression-15969]
	_ = x[ErrSortBadExpression-15973]
	_ = x[ErrSortBadValue-15974]
	_ = x[ErrSortBadOrder-15975]
	_ = x[ErrSortMissingKey-15976]
	_ = x[ErrStageUnwindWrongType-15981]
	_ = x[ErrPathContainsEmptyElement-15998]
	_ = x[ErrFieldPathInvalidName-16410]
	_ = x[ErrGroupInvalidFieldPath-16872]
	_ = x[ErrGroupUndefinedVariable-17276]
	_ = x[ErrInvalidArg-28667]
	_ = x[ErrSliceFirstArg-28724]
	_ = x[ErrStageUnwindNoPath-28812]
	_ = x[ErrStageUnwindNoPrefix-28818]
	_ = x[ErrProjectionInEx-31253]
	_ = x[ErrProjectionExIn-31254]
	_ = x[ErrStageCountNonString-40156]
	_ = x[ErrStageCountNonEmptyString-40157]
	_ = x[ErrStageCountBadPrefix-40158]
	_ = x[ErrStageCountBadValue-40160]
	_ = x[ErrStageGroupUnaryOperator-40237]
	_ = x[ErrStageGroupMultipleAccumulator-40238]
	_ = x[ErrStageGroupInvalidAccumulator-40234]
	_ = x[ErrStageInvalid-40323]
	_ = x[ErrEmptyFieldPath-40352]
	_ = x[ErrMissingField-40414]
	_ = x[ErrFailedToParseInput-40415]
	_ = x[ErrCollStatsIsNotFirstStage-40415]
	_ = x[ErrFreeMonitoringDisabled-50840]
	_ = x[ErrValueNegative-51024]
	_ = x[ErrRegexOptions-51075]
	_ = x[ErrRegexMissingParen-51091]
	_ = x[ErrBadRegexOption-51108]
	_ = x[ErrDuplicateField-4822819]
	_ = x[ErrStageSkipBadValue-5107200]
	_ = x[ErrStageLimitInvalidArg-5107201]
	_ = x[ErrStageCollStatsInvalidArg-5447000]
}

const _ErrorCode_name = "UnsetInternalErrorBadValueFailedToParseTypeMismatchNamespaceNotFoundIndexNotFoundUnsuitableValueTypeConflictingUpdateOperatorsCursorNotFoundNamespaceExistsDollarPrefixedFieldNameInvalidIDEmptyNameCommandNotFoundImmutableFieldCannotCreateIndexInvalidOptionsInvalidNamespaceIndexOptionsConflictIndexKeySpecsConflictOperationFailedDocumentValidationFailureNotImplementedMechanismUnavailableLocation11000Location15947Location15948Location15955Location15958Location15959Location15969Location15973Location15974Location15975Location15976Location15981Location15998Location16410Location16872Location17276Location28667Location28724Location28812Location28818Location31253Location31254Location40156Location40157Location40158Location40160Location40234Location40237Location40238Location40323Location40352Location40414Location40415Location50840Location51024Location51075Location51091Location51108Location4822819Location5107200Location5107201Location5447000"

var _ErrorCode_map = map[ErrorCode]string{
	0:       _ErrorCode_name[0:5],
	1:       _ErrorCode_name[5:18],
	2:       _ErrorCode_name[18:26],
	9:       _ErrorCode_name[26:39],
	14:      _ErrorCode_name[39:51],
	26:      _ErrorCode_name[51:68],
	27:      _ErrorCode_name[68:81],
	28:      _ErrorCode_name[81:100],
	40:      _ErrorCode_name[100:126],
	43:      _ErrorCode_name[126:140],
	48:      _ErrorCode_name[140:155],
	52:      _ErrorCode_name[155:178],
	53:      _ErrorCode_name[178:187],
	56:      _ErrorCode_name[187:196],
	59:      _ErrorCode_name[196:211],
	66:      _ErrorCode_name[211:225],
	67:      _ErrorCode_name[225:242],
	72:      _ErrorCode_name[242:256],
	73:      _ErrorCode_name[256:272],
	85:      _ErrorCode_name[272:292],
	86:      _ErrorCode_name[292:313],
	96:      _ErrorCode_name[313:328],
	121:     _ErrorCode_name[328:353],
	238:     _ErrorCode_name[353:367],
	334:     _ErrorCode_name[367:387],
	11000:   _ErrorCode_name[387:400],
	15947:   _ErrorCode_name[400:413],
	15948:   _ErrorCode_name[413:426],
	15955:   _ErrorCode_name[426:439],
	15958:   _ErrorCode_name[439:452],
	15959:   _ErrorCode_name[452:465],
	15969:   _ErrorCode_name[465:478],
	15973:   _ErrorCode_name[478:491],
	15974:   _ErrorCode_name[491:504],
	15975:   _ErrorCode_name[504:517],
	15976:   _ErrorCode_name[517:530],
	15981:   _ErrorCode_name[530:543],
	15998:   _ErrorCode_name[543:556],
	16410:   _ErrorCode_name[556:569],
	16872:   _ErrorCode_name[569:582],
	17276:   _ErrorCode_name[582:595],
	28667:   _ErrorCode_name[595:608],
	28724:   _ErrorCode_name[608:621],
	28812:   _ErrorCode_name[621:634],
	28818:   _ErrorCode_name[634:647],
	31253:   _ErrorCode_name[647:660],
	31254:   _ErrorCode_name[660:673],
	40156:   _ErrorCode_name[673:686],
	40157:   _ErrorCode_name[686:699],
	40158:   _ErrorCode_name[699:712],
	40160:   _ErrorCode_name[712:725],
	40234:   _ErrorCode_name[725:738],
	40237:   _ErrorCode_name[738:751],
	40238:   _ErrorCode_name[751:764],
	40323:   _ErrorCode_name[764:777],
	40352:   _ErrorCode_name[777:790],
	40414:   _ErrorCode_name[790:803],
	40415:   _ErrorCode_name[803:816],
	50840:   _ErrorCode_name[816:829],
	51024:   _ErrorCode_name[829:842],
	51075:   _ErrorCode_name[842:855],
	51091:   _ErrorCode_name[855:868],
	51108:   _ErrorCode_name[868:881],
	4822819: _ErrorCode_name[881:896],
	5107200: _ErrorCode_name[896:911],
	5107201: _ErrorCode_name[911:926],
	5447000: _ErrorCode_name[926:941],
}

func (i ErrorCode) String() string {
	if str, ok := _ErrorCode_map[i]; ok {
		return str
	}
	return "ErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
}