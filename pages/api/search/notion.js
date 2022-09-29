export default function notion(req, res){

    const { Client } = require("@notionhq/client")

    const notion = new Client({
        auth: "",
    });

    (async () => {
        const response = await notion.search({
        query: 'エンジニア',
        sort: {
            direction: 'ascending',
            timestamp: 'last_edited_time',
        },
        });
        res.status(200).json(response)
    })();
}