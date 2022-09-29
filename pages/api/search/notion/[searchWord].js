export default function notion(req, res){

    const searchWord = req.query.searchWord

    const { Client } = require("@notionhq/client")

    const notion = new Client({
        auth: "secret_tG92Vh2nvwyrB7JKJVFfLRAGeTfb4HyW8iIrvm8xuUj",
    });

    (async () => {
        const response = await notion.search({
        query: searchWord,
        sort: {
            direction: 'ascending',
            timestamp: 'last_edited_time',
        },
        });

        let results = []
        for (let i=0;i<response.results.length;i++) {
            /* parse of
            "properties": {
                "サービスの機能": {
                    "id": "%5Ddff",
                    "type": "rich_text",
                    "rich_text": [
                        {
                            "type": "text",
                            "text": {
                                "content": "コンテナー化されたサービスを実行する VM のクラスター管理。",
                                "link": null
                            },
                            "annotations": {
                                "bold": false,
                                "italic": false,
                                "strikethrough": false,
                                "underline": false,
                                "code": false,
                                "color": "default"
                            },
                            "plain_text": "コンテナー化されたサービスを実行する VM のクラスター管理。",
                            "href": null
                        }
                    ]
                },
                "サービス名": {
                    "id": "title",
                    "type": "title",
                    "title": [
                        {
                            "type": "text",
                            "text": {
                                "content": "Azure Kubernetes Service",
                                "link": null
                            },
                            "annotations": {
                                "bold": false,
                                "italic": false,
                                "strikethrough": false,
                                "underline": false,
                                "code": false,
                                "color": "default"
                            },
                            "plain_text": "Azure Kubernetes Service",
                            "href": null
                        }
                    ]
                }
            }
            */
            
            let context = ""
            const dataobj = response.results[i].properties

            for (let k in dataobj) {
                if(Object.keys(dataobj[k]).indexOf("title")==2){
                    // キーの中にtitleキーが含まれる場合
                    for(let j in dataobj[k].title){
                        context += dataobj[k].title[j].plain_text
                    }
                }else if(Object.keys(dataobj[k]).indexOf("rich_text")==2){
                    // キーの中にrich_textキーが含まれる場合
                    for(let j in dataobj[k].rich_text){
                        context += dataobj[k].rich_text[j].plain_text
                    }
                }
            }
            results.push({
                "context": context,
                "url": response.results[i].url,
            })
        }
        res.status(200).json({results})
    })();
}