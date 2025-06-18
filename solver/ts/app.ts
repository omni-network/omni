import express, { Request, Response } from 'express';
import bodyParser from 'body-parser';

const app = express();
app.use(bodyParser.json());

type Req = {
    from: string;
    amount: number;
};

type Resp = {
    to: string;
    amount: number;
};

app.post('/', (req: Request, res: Response) => {
    console.log('body:', req.body);

    const request: Req = req.body;

    // log request
    console.log('Request:', request);

    const response: Resp = {
        to: request.from + "Z",
        amount: request.amount + 1
    }
    console.log('Response:', response);

    res.json( response );
});


const port = 8000;
app.listen(port, () => {
    console.log(`Server is running on port ${port}`);
});
