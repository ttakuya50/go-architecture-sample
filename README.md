
# go-architecture-sample

## アーキテクチャ説明
- 別記事にします

## 動作確認
```shell
make docker_up // ローカルでDB起動
make run // API起動
curl -X POST "127.0.0.1:8888/user" -d '{"name":"aaaa"}'
curl -X DELETE "127.0.0.1:8888/user" -d '{"user_id":12345}'
curl -X POST "127.0.0.1:8888/user/list" -d '{"user_id":2810422652536671453,"title":"テスト1"}'
```

## 参考資料
- [エリック・エヴァンスのドメイン駆動設計 (IT Architects’Archive ソフトウェア開発の実践)](https://www.amazon.co.jp/%E3%82%A8%E3%83%AA%E3%83%83%E3%82%AF%E3%83%BB%E3%82%A8%E3%83%B4%E3%82%A1%E3%83%B3%E3%82%B9%E3%81%AE%E3%83%89%E3%83%A1%E3%82%A4%E3%83%B3%E9%A7%86%E5%8B%95%E8%A8%AD%E8%A8%88-Architects%E2%80%99Archive-%E3%82%BD%E3%83%95%E3%83%88%E3%82%A6%E3%82%A7%E3%82%A2%E9%96%8B%E7%99%BA%E3%81%AE%E5%AE%9F%E8%B7%B5-%E3%82%A8%E3%83%AA%E3%83%83%E3%82%AF%E3%83%BB%E3%82%A8%E3%83%B4%E3%82%A1%E3%83%B3%E3%82%B9/dp/4798121967/ref=asc_df_4798121967/?tag=jpgo-22&linkCode=df0&hvadid=295719984664&hvpos=&hvnetw=g&hvrand=12144184347501724598&hvpone=&hvptwo=&hvqmt=&hvdev=c&hvdvcmdl=&hvlocint=&hvlocphy=1009308&hvtargid=pla-525481409888&psc=1&th=1&psc=1)
- [実践ドメイン駆動設計 (Object Oriented SELECTION)](https://www.amazon.co.jp/%E5%AE%9F%E8%B7%B5%E3%83%89%E3%83%A1%E3%82%A4%E3%83%B3%E9%A7%86%E5%8B%95%E8%A8%AD%E8%A8%88-Object-Oriented-SELECTION-%E3%83%B4%E3%82%A1%E3%83%BC%E3%83%B3%E3%83%BB%E3%83%B4%E3%82%A1%E3%83%BC%E3%83%8E%E3%83%B3/dp/479813161X/ref=asc_df_479813161X/?tag=jpgo-22&linkCode=df0&hvadid=295706574430&hvpos=&hvnetw=g&hvrand=12144184347501724598&hvpone=&hvptwo=&hvqmt=&hvdev=c&hvdvcmdl=&hvlocint=&hvlocphy=1009308&hvtargid=pla-525785668462&psc=1&th=1&psc=1)
- [現場で役立つシステム設計の原則 ~変更を楽で安全にするオブジェクト指向の実践技法](https://www.amazon.co.jp/%E7%8F%BE%E5%A0%B4%E3%81%A7%E5%BD%B9%E7%AB%8B%E3%81%A4%E3%82%B7%E3%82%B9%E3%83%86%E3%83%A0%E8%A8%AD%E8%A8%88%E3%81%AE%E5%8E%9F%E5%89%87-%E5%A4%89%E6%9B%B4%E3%82%92%E6%A5%BD%E3%81%A7%E5%AE%89%E5%85%A8%E3%81%AB%E3%81%99%E3%82%8B%E3%82%AA%E3%83%96%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88%E6%8C%87%E5%90%91%E3%81%AE%E5%AE%9F%E8%B7%B5%E6%8A%80%E6%B3%95-%E5%A2%97%E7%94%B0-%E4%BA%A8/dp/477419087X/ref=asc_df_477419087X/?tag=jpgo-22&linkCode=df0&hvadid=295686767484&hvpos=&hvnetw=g&hvrand=8625275969429694189&hvpone=&hvptwo=&hvqmt=&hvdev=c&hvdvcmdl=&hvlocint=&hvlocphy=1009308&hvtargid=pla-526272648873&psc=1&th=1&psc=1)
