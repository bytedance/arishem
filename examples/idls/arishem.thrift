// BoolConst or NumConst or StrConst
struct Const {
    1: optional bool BoolConst
    2: optional double NumConst
    3: optional string StrConst
}

// (Lhs and Rhs) or (ParamList)
struct MathExpr {
    1: required string OpMath
    2: optional Expr Lhs
    3: optional Expr Rhs
    4: optional list<Expr> ParamList
}

// ParamMap or ParamList
struct FuncExpr {
    1: required string FuncName
    2: optional map<string,Expr> ParamMap
    3: optional list<Expr> ParamList
}

struct FeatureExpr {
    1: required string FeaturePath
    2: optional map<string,Expr> BuiltinParam
}

struct Expr {
    1: optional list<Expr> ListExpr
    2: optional map<string,Expr> MapExpr
    3: optional Const Const
    4: optional list<Const> ConstList
    5: optional MathExpr MathExpr
    6: optional FuncExpr FuncExpr
    7: optional string VarExpr
    8: optional FeatureExpr FeatureExpr
}

// ParamMap or ParamList
struct ActionAim {
    1: required string ActionName
    2: optional map<string,Expr> ParamMap
    3: optional list<Expr> ParamList
}

struct Condition {
    1: required string Operator
    2: required Expr Lhs
    3: required Expr Rhs
}

// ConditionGroups or Conditions
struct ConditionGroup {
    1: required string OpLogic
    2: optional list<ConditionGroup> ConditionGroups
    3: optional list<Condition> Conditions
}

struct RuleExample {
    1: required string Name
    2: required ConditionGroup ConditionExpr
    3: optional ActionAim ActionAim
    4: optional list<ActionAim> ActionListAim
    5: optional Expr ExprAim
}