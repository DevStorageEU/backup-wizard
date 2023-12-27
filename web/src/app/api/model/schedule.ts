/**
 * Backup Wizard API
 * An API to communicate with the backup wizard daemon
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { RetentionKind } from './retentionKind';
import { Weekday } from './weekday';


export interface Schedule { 
    kind: RetentionKind;
    weekday: Weekday;
    time: string;
}
export namespace Schedule {
}

