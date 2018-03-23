/**
 * GatewayWebsocketApi
 * Websocket API Spec for Gateway
 *
 * OpenAPI spec version: 0.0.1
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


export interface WebSocketResponseHeader {
    responseType: WebSocketResponseHeader.ResponseTypeEnum;
    error?: Error;
}
export namespace WebSocketResponseHeader {
    export type ResponseTypeEnum = 'Error' | 'UpdateTotalAccessCount' | 'UpdateConnectionCount' | 'UpdateMasterIdentifier' | 'UpdateMasterNodeCount';
    export const ResponseTypeEnum = {
        Error: 'Error' as ResponseTypeEnum,
        UpdateTotalAccessCount: 'UpdateTotalAccessCount' as ResponseTypeEnum,
        UpdateConnectionCount: 'UpdateConnectionCount' as ResponseTypeEnum,
        UpdateMasterIdentifier: 'UpdateMasterIdentifier' as ResponseTypeEnum,
        UpdateMasterNodeCount: 'UpdateMasterNodeCount' as ResponseTypeEnum
    }
}
