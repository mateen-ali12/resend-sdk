import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { SuppressionBatch, SuppressionBatchCreateData } from '../ResendTypes';
declare class SuppressionBatchEntity extends ResendEntityBase<SuppressionBatch> {
    constructor(client: ResendSDK, entopts: any);
    make(this: SuppressionBatchEntity): SuppressionBatchEntity;
    create(this: any, reqdata?: SuppressionBatchCreateData, ctrl?: Control): Promise<SuppressionBatchEntity>;
}
export { SuppressionBatchEntity };
