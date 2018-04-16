package config

import (
    "io/ioutil"
    "encoding/json"
)

// 構造体定義
type DBConfig struct {
    DBMS string  `json:"dbms"`
    User string  `json:"user"`
    Pass  string  `json:"pass"`
    Protcol  string  `json:"protcol"`
    DBName  string  `json:"dbname"`
}

// configファイルを読み込み構造体へ割当
func ReadDBConfig(filename string) (*DBConfig, error) {

    // 構造体を定義
    c := new(DBConfig)

    // 設定ファイルを読み込む
    jsonString, err := ioutil.ReadFile(filename)
    if err != nil {
        // エラー
        return c, err
    }

    // 設定
    err = json.Unmarshal(jsonString, c)
    if err != nil {
        // エラー
        return c, err
    }

    // 正常
    return c, nil
}
