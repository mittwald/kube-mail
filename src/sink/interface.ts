import {SMTPServerSession} from "smtp-server";
import {ParsedMail} from "mailparser";

export interface Message {
    envelope: {
        mailFrom: string;
        rcptTo: string[];
    };

    date: Date;
    mail: ParsedMail;
    remoteAddress?: string;
}

export interface Parser {
    parseMessage(session: SMTPServerSession, data: Buffer): Promise<Message>
}
