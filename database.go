package main

import (
	"fmt"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

func (d *DataBase) OpenDatabase(name string) error {
	db, err := bbolt.Open(fmt.Sprintf("%s.db", name), 0644, nil)
	if err != nil {
		return fmt.Errorf("OpenDatabase: %v", err)
	}
	d.DB = db
	return nil
}

func (d *DataBase) CloseDatabase() error {
	err := d.DB.Close()
	if err != nil {
		return fmt.Errorf("CloseDatabase: %v", err)
	}
	return nil
}

func (d *DataBase) WriteTestData(
	dataVolume int,
) {
	d.DB.Update(
		func(tx *bbolt.Tx) error {
			tx.DeleteBucket([]byte("TestBacket"))
			// 删除名为 TestBacket 的存储桶，如果有的话
			bucket, err := tx.CreateBucket([]byte("TestBacket"))
			if err != nil {
				return fmt.Errorf("WriteTestData: %v", err)
			}
			// 创建存储桶，名为 TestBacket
			for i := 0; i < dataVolume; i++ {
				uuidForKey := uuid.New()
				uuidForValue := uuid.New()
				err = bucket.Put(uuidForKey[:], uuidForValue[:])
				if err != nil {
					return fmt.Errorf("WriteTestData: %v", err)
				}
			}
			// 随机生成数据，然后写入到数据库中
			return nil
			// 返回值
		},
	)
}

func (d *DataBase) AppendTestData(
	dataVolume int,
) {
	d.DB.Update(
		func(tx *bbolt.Tx) error {
			bucket := tx.Bucket([]byte("TestBacket"))
			// 取得名为 TestBacket 的存储桶
			for i := 0; i < dataVolume; i++ {
				uuidForKey := uuid.New()
				uuidForValue := uuid.New()
				err := bucket.Put(uuidForKey[:], uuidForValue[:])
				if err != nil {
					return fmt.Errorf("AppendTestData: %v", err)
				}
			}
			// 随机生成数据，然后写入到数据库中
			return nil
			// 返回值
		},
	)
}

func (d *DataBase) GetTestData() {
	d.DataReceived = make(map[uuid.UUID]uuid.UUID)
	d.DB.View(
		func(tx *bbolt.Tx) error {
			bucket := tx.Bucket([]byte("TestBacket"))
			// 取得名为 TestBacket 的存储桶
			err := bucket.ForEach(
				func(k, v []byte) error {
					d.DataReceived[uuid.UUID(k)] = uuid.UUID(v)
					return nil
				},
			)
			if err != nil {
				return fmt.Errorf("GetTestData: %v", err)
			}
			// 从存储桶获取数据
			return nil
			// 返回值
		},
	)
}
