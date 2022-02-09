package control

/*
	java 中的switch-case语句有两种实现方式：
		1. 如果case可以被编码成一个索引表，则实现成tableswitch指令；
		2. 否则实现成lookupswitch指令。
	比如这个就会被编译成tableswitch：

	int chooseNear(int i ){
		switch (i) {
			case 0 : return 0;
			case 1 : return 1;
			case 2 : return 2;
			default: return -1;
		}
	}

	这个则会被编码成lookupswitch
	int chooseNear(int i ){
		switch (i) {
			case -100 : return -1;
			case 0 : return 0;
			case 100 : return 1;
			default: return -1;
		}
	}
*/

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}
