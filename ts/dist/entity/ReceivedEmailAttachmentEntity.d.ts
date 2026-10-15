import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { ReceivedEmailAttachment, ReceivedEmailAttachmentLoadMatch, ReceivedEmailAttachmentListMatch } from '../ResendTypes';
declare class ReceivedEmailAttachmentEntity extends ResendEntityBase<ReceivedEmailAttachment> {
    constructor(client: ResendSDK, entopts: any);
    make(this: ReceivedEmailAttachmentEntity): ReceivedEmailAttachmentEntity;
    load(this: any, reqmatch?: ReceivedEmailAttachmentLoadMatch, ctrl?: Control): Promise<ReceivedEmailAttachmentEntity>;
    list(this: any, reqmatch?: ReceivedEmailAttachmentListMatch, ctrl?: Control): Promise<ReceivedEmailAttachmentEntity[]>;
}
export { ReceivedEmailAttachmentEntity };
