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
    EntTeacherEdges,
    EntTeacherEdgesFromJSON,
    EntTeacherEdgesFromJSONTyped,
    EntTeacherEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntTeacher
 */
export interface EntTeacher {
    /**
     * 
     * @type {EntTeacherEdges}
     * @memberof EntTeacher
     */
    edges?: EntTeacherEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntTeacher
     */
    id?: number;
    /**
     * TeacherEmail holds the value of the "teacher_email" field.
     * @type {string}
     * @memberof EntTeacher
     */
    teacherEmail?: string;
    /**
     * TeacherName holds the value of the "teacher_name" field.
     * @type {string}
     * @memberof EntTeacher
     */
    teacherName?: string;
}

export function EntTeacherFromJSON(json: any): EntTeacher {
    return EntTeacherFromJSONTyped(json, false);
}

export function EntTeacherFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntTeacher {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntTeacherEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'teacherEmail': !exists(json, 'teacher_email') ? undefined : json['teacher_email'],
        'teacherName': !exists(json, 'teacher_name') ? undefined : json['teacher_name'],
    };
}

export function EntTeacherToJSON(value?: EntTeacher | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntTeacherEdgesToJSON(value.edges),
        'id': value.id,
        'teacher_email': value.teacherEmail,
        'teacher_name': value.teacherName,
    };
}

