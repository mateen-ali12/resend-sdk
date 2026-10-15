import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { EmailAttachment, EmailAttachmentLoadMatch, EmailAttachmentListMatch } from '../ResendTypes';
declare class EmailAttachmentEntity extends ResendEntityBase<EmailAttachment> {
    constructor(client: ResendSDK, entopts: any);
    make(this: EmailAttachmentEntity): EmailAttachmentEntity;
    load(this: any, reqmatch?: EmailAttachmentLoadMatch, ctrl?: Control): Promise<EmailAttachmentEntity>;
    list(this: any, reqmatch?: EmailAttachmentListMatch, ctrl?: Control): Promise<EmailAttachmentEntity[]>;
}
export { EmailAttachmentEntity };
