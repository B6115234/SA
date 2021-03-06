/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Project API Course Item
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntSubjectEdges,
    EntSubjectEdgesFromJSON,
    EntSubjectEdgesFromJSONTyped,
    EntSubjectEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntSubject
 */
export interface EntSubject {
    /**
     * 
     * @type {EntSubjectEdges}
     * @memberof EntSubject
     */
    edges?: EntSubjectEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntSubject
     */
    id?: number;
    /**
     * SubjectName holds the value of the "subject_name" field.
     * @type {string}
     * @memberof EntSubject
     */
    subjectName?: string;
}

export function EntSubjectFromJSON(json: any): EntSubject {
    return EntSubjectFromJSONTyped(json, false);
}

export function EntSubjectFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntSubject {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntSubjectEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'subjectName': !exists(json, 'subject_name') ? undefined : json['subject_name'],
    };
}

export function EntSubjectToJSON(value?: EntSubject | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntSubjectEdgesToJSON(value.edges),
        'id': value.id,
        'subject_name': value.subjectName,
    };
}


