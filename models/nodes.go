package models

import (
	"blog/define"
	"github.com/jinzhu/gorm"
)

type Nodes []*Node

type Node struct {
	ID int64 `gorm:"primary_key" json:"id"`
	ParentID int64 `json:"parent_id"`
	Type define.NodeType `json:"type"`
	FilePath string `json:"file_path"`
	CategoryID int64 `json:"category_id"`
	Name string `json:"name"`
	Discription string `json:"discription"`
	Children Nodes `json:"children"`
}

func (*Node) TableName() string{
	return "node"
}

//新增节点
func (n *Node) Insert(db *gorm.DB) error{
	return db.Table(n.TableName()).Create(n).Error
}

//删除节点
func (n *Node) Delete(db *gorm.DB) (int64,error){
	db = db.Table(n.TableName()).Where("id = ?",n.ID).Delete(nil)
	return db.RowsAffected,db.Error
}

//修改节点
func (n *Node) Update(db *gorm.DB,updateParams map[string]interface{}) (int64,error){
	db  = db.Table(n.TableName()).Where("id  = ?" ,n.ID).Update(updateParams)
	return db.RowsAffected,db.Error
}

//查找所有节点
func (n *Node) List(db *gorm.DB,categoryId int64)(list Nodes,err error){
	err = db.Table(n.TableName()).Where("category_id = ?",categoryId).Find(&list).Error

	if err ==nil{
		list = list.ToTree()
	}

	return
}

//树形数据
func (list Nodes) ToTree() Nodes{

	_map := make(map[int64]*Node,len(list))

	res := make(Nodes,0)

	for _,val := range list{
		_map[val.ID] = val
		if val.ParentID >0{
			_map[val.ParentID].Children = append(_map[val.ParentID].Children,val)
		}else{
			res = append(res,val)
		}

	}

	return res

}




