package transfer

import (
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// ConvertToAny 将源对象转换为 protobuf 的 Any 类型
// data: 源对象指针
// target: 目标 protobuf 消息类型的实例（用于反射创建对象）
func ConvertToAny(data any, target proto.Message) (*anypb.Any, error) {
	// 创建目标类型的新实例
	// targetClone := proto.Clone(target)

	// 使用 copier 复制字段
	if err := copier.Copy(target, data); err != nil {
		return nil, err
	}

	// 将 proto.Message 转换为 Any 类型
	return anypb.New(target)
}
