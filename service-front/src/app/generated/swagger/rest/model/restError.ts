/**
 * GatewayRestApi
 * REST API Spec for Gateway
 *
 * OpenAPI spec version: 0.0.1
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


export interface RestError {
    timestamp?: string;
    code?: number;
    message?: string;
    type?: RestError.TypeEnum;
}
export namespace RestError {
    export type TypeEnum = 'InvalidSession' | 'InternalServer' | 'BadFilterRequest' | 'RecordDoesNotxist';
    export const TypeEnum = {
        InvalidSession: 'InvalidSession' as TypeEnum,
        InternalServer: 'InternalServer' as TypeEnum,
        BadFilterRequest: 'BadFilterRequest' as TypeEnum,
        RecordDoesNotxist: 'RecordDoesNotxist' as TypeEnum
    }
}
