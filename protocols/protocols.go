package protocols

// 枚举协议号
const (
	/**********************  登录  **********************/
	// 登录验证
	P_Login_Validate = 1000001
	// 登录请求角色列表
	P_Login_RequestRole = 1000002
	// 登录请求创建角色
	P_Login_CreateRole = 1000003
	// 登录请求选择角色
	P_Login_ChooseRole = 1000004
	// 登录验证在线
	P_Login_ValidateOnline = 1000005
	// 登录验证Ping
	P_Login_Ping = 1000006

	/**********************  角色  **********************/
	// 角色-角色进入游戏
	P_Role_RoleEnterLogic = 2000001

	// 角色-同步角色数据
	P_Role_SynRoleData = 2000002

	// 角色-同步角色属性
	P_Role_SynRoleAttrValue = 2000003

	// 角色-同步角色信息
	P_Role_SynRoleInformationData = 2000004

	// 角色-同步角色背包数据
	P_Role_SynItemData = 2000006

	// 角色-同步角色邮件数据
	P_Role_SynMailData = 2000007

	// 角色-同步角色任务数据
	P_Role_SynTaskData = 2000009

	// 角色-同步战斗阵容数据
	P_Role_SynBattleArrayData = 2000011

	// 角色-英雄升级操作
	P_Role_HeroLevelUp = 2000024

	// 角色-设置默认战斗阵容
	P_Role_BattleArraySetDefine = 2000025

	// 角色-设置战斗阵容
	P_Role_BattleArrayUp = 2000026

	// 角色-角色对战结算数据
	P_Role_FightBalance = 2000030

	// 角色-设置引导步骤
	P_Role_SetGuide = 2000034

	// 角色-角色简要信息
	P_Role_GetRoleSimpleInfo = 2000039

	// 角色-同步外挂数据
	P_Role_SynIndulge = 2000052

	// 角色-角色总看广告数据
	P_Role_TotalWatchADBoxData = 2000054

	// 角色-同步抽奖数据
	P_Role_SyncDrawPrize = 2000060

	// 角色-抽奖
	P_Role_DrawPrize = 2000061

	// 角色-同步角色任务额外数据
	P_Role_SynTaskExtraData = 2000071

	// 角色-同步角色条件数据
	P_Role_SynCondShareData = 2000105

	// 角色-角色花费数据
	P_Role_Cost_Get = 2000113

	// 角色-同步角色开关数据
	P_Role_OnOffDataInfo = 2000121

	// 角色-同步角色花费数据
	P_Role_SyncCostGet = 2000131

	// 角色-修改角色战车皮肤
	P_Role_Car_Skin_Change = 2000138

	// 角色-发送卡组血量到服务器
	P_Role_SendCarHpToServer = 2000148

	// 角色-修改角色英雄皮肤
	P_Role_HeroChangeSkin = 2000151

	// 角色-修改角色战斗阵容名称
	P_Role_SetBattleArrayName = 2000153

	// 角色-同步角色章节数据
	P_Role_SynChapterData = 2000155

	// 角色-战车链接
	P_Role_CarLink = 2000163

	/**********************  活动  **********************/
	// 活动-同步角色所有活动数据
	P_Activity_SynAllActivityData = 4000001

	// 活动-获取颁奖日数据
	P_Activity_GetDayPrizeData = 4000014

	// 活动-获取黄金联赛数据
	P_Activity_GetGoldenLeagueData = 4000016

	// 活动-同步角色试炼场数据
	P_Activity_SyncEatChickenData = 4000047

	// 活动-同步角色公会战数据
	P_Activity_SyncSociatyWarData = 4000050

	// 活动-获取角色大航海数据
	P_Activity_GetGreatSailingData = 4000052

	// 活动-同步角色大航海数据
	P_Activity_SyncGreatSailingData = 4000053

	// 活动-大航海刷新卡组
	P_Activity_GreatSailingRefleshCard = 4000055

	// 活动-同步角色天空之城数据
	P_Activity_SyncSkyCastleData = 4000057

	// 活动-同步角色寒冰堡数据
	P_Activity_SyncWeekCooperationData = 4000060

	// 活动-同步角色周年庆数据
	P_Activity_SyncAnniversaryTrialData = 4000063

	// 活动-同步角色回归数据
	P_Activity_SyncReturnBackData = 4000065

	// 活动-同步角色雾隐数据
	P_Activity_SyncFogHiddenData = 4000067

	// 活动-机械迷城数据
	P_Activity_SyncMachinariumData = 4000069

	/**********************  聊天  **********************/
	// 聊天请求
	P_Chat_ToClient = 6000001
	// 战斗匹配房间关闭
	P_Chat_CloseFightRoom = 6000003

	// 同步朋友数据
	P_Friend_SynLoginData = 7000001

	/**********************  战斗  **********************/
	// 战斗匹配
	P_Match_Fight = 9000001
	// 战斗匹配取消
	P_Match_Cancel = 9000002
	// 战斗匹配结果
	P_Match_Result = 9000003
	// 战斗匹配竞争战斗
	P_Match_Duel_Fight = 9000004
	// 战斗匹配对战取消
	P_Match_Duel_Cancel = 9000005

	// 对战登录
	P_Fight_Role_Login = 10000002
	// 对战加载准备
	P_Fight_Loading_Ready = 10000003
	// 对战开始
	P_Fight_FightStart = 10000004
	// 对战结束
	P_Fight_FightEnd = 10000005
	// 怪物血量同步
	P_Fight_Monster_Blood_SYNC = 10000007
	// 战斗对战结束提交（多人，结束）
	P_Fight_Report_Result_To_Fight = 10000014
	// 战斗匹配对战结束(单人，结束)
	P_Fight_Report_Result_To_Logic = 10000015
	// 提交对战每阶段逻辑数据（多人，过程）
	P_Fight_Report_Phase_Result_To_Fight = 10000020
	// 提交对战每阶段逻辑数据（单人，过程）
	P_Fight_Report_Phase_Result_To_Logic = 10000021
	// 更新英雄
	P_Fight_Update_Hero_SYNC = 10000101
	// 战斗银币同步
	P_Fight_Silver_SYNC = 10000104
	// 战斗卖出英雄同步
	P_Fight_Sell_Hero_SYNC = 10000105
	// 战斗英雄属性同步
	P_Fight_Hero_Attr_SYNC = 10000107
	// 卡牌刷新次数同步
	P_Fight_Refresh_Card_Count_SYNC = 10000108
	// 操作装备同步
	P_Fight_Operate_Equip_SYNC = 10000109

	/**********************  联盟  **********************/

	// 同步联盟数据
	P_Sociaty_SynData = 12000001
	// 获取机械迷城数据
	P_Sociaty_RoleGetMachinariumData = 12000033
	// 同步机械迷城数据
	P_Sociaty_SyncMachinariumData = 12000034
	// 机械迷城选择卡组
	P_Sociaty_RoleMachinariumSelectCard = 12000037

	/**********************  其他  **********************/
	// 对战网络到主逻辑服务
	P_Network_Fight_To_Logic = 99999999
)
