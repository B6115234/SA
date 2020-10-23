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


import * as runtime from '../runtime';
import {
    EntCourse,
    EntCourseFromJSON,
    EntCourseToJSON,
    EntCourseItem,
    EntCourseItemFromJSON,
    EntCourseItemToJSON,
    EntSubject,
    EntSubjectFromJSON,
    EntSubjectToJSON,
    EntSubjectType,
    EntSubjectTypeFromJSON,
    EntSubjectTypeToJSON,
    EntTeacher,
    EntTeacherFromJSON,
    EntTeacherToJSON,
} from '../models';

export interface CreateCourseRequest {
    course: EntCourse;
}

export interface CreateCourseitemRequest {
    courseitem: EntCourseItem;
}

export interface CreateSubjectRequest {
    subject: EntSubject;
}

export interface CreateSubjectTypeRequest {
    subjectType: EntCourse;
}

export interface CreateTeacherRequest {
    teacher: EntCourse;
}

export interface DeleteCourseitemRequest {
    id: number;
}

export interface DeleteTeacherRequest {
    id: number;
}

export interface GetCourseRequest {
    id: number;
}

export interface GetCourseitemRequest {
    id: number;
}

export interface GetSubjectRequest {
    id: number;
}

export interface GetSubjectTypeRequest {
    id: number;
}

export interface GetTeacherRequest {
    id: number;
}

export interface ListCourseRequest {
    limit?: number;
    offset?: number;
}

export interface ListCourseitemRequest {
    limit?: number;
    offset?: number;
}

export interface ListSubjectRequest {
    limit?: number;
    offset?: number;
}

export interface ListSubjectTypeRequest {
    limit?: number;
    offset?: number;
}

export interface ListTeacherRequest {
    limit?: number;
    offset?: number;
}

export interface UpdateCourseitemRequest {
    id: number;
    courseitem: EntCourseItem;
}

export interface UpdateTeacherRequest {
    id: number;
    teacher: EntTeacher;
}

/**
 * 
 */
export class DefaultApi extends runtime.BaseAPI {

    /**
     * Create Course
     * Create Course
     */
    async createCourseRaw(requestParameters: CreateCourseRequest): Promise<runtime.ApiResponse<EntCourse>> {
        if (requestParameters.course === null || requestParameters.course === undefined) {
            throw new runtime.RequiredError('course','Required parameter requestParameters.course was null or undefined when calling createCourse.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/Course`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntCourseToJSON(requestParameters.course),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntCourseFromJSON(jsonValue));
    }

    /**
     * Create Course
     * Create Course
     */
    async createCourse(requestParameters: CreateCourseRequest): Promise<EntCourse> {
        const response = await this.createCourseRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create courseitem
     * Create courseitem
     */
    async createCourseitemRaw(requestParameters: CreateCourseitemRequest): Promise<runtime.ApiResponse<EntCourseItem>> {
        if (requestParameters.courseitem === null || requestParameters.courseitem === undefined) {
            throw new runtime.RequiredError('courseitem','Required parameter requestParameters.courseitem was null or undefined when calling createCourseitem.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/CourseItems`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntCourseItemToJSON(requestParameters.courseitem),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntCourseItemFromJSON(jsonValue));
    }

    /**
     * Create courseitem
     * Create courseitem
     */
    async createCourseitem(requestParameters: CreateCourseitemRequest): Promise<EntCourseItem> {
        const response = await this.createCourseitemRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create Subject
     * Create Subject
     */
    async createSubjectRaw(requestParameters: CreateSubjectRequest): Promise<runtime.ApiResponse<EntSubject>> {
        if (requestParameters.subject === null || requestParameters.subject === undefined) {
            throw new runtime.RequiredError('subject','Required parameter requestParameters.subject was null or undefined when calling createSubject.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/Subjects`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntSubjectToJSON(requestParameters.subject),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntSubjectFromJSON(jsonValue));
    }

    /**
     * Create Subject
     * Create Subject
     */
    async createSubject(requestParameters: CreateSubjectRequest): Promise<EntSubject> {
        const response = await this.createSubjectRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create SubjectType
     * Create SubjectType
     */
    async createSubjectTypeRaw(requestParameters: CreateSubjectTypeRequest): Promise<runtime.ApiResponse<EntSubjectType>> {
        if (requestParameters.subjectType === null || requestParameters.subjectType === undefined) {
            throw new runtime.RequiredError('subjectType','Required parameter requestParameters.subjectType was null or undefined when calling createSubjectType.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/SubjectTypes`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntCourseToJSON(requestParameters.subjectType),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntSubjectTypeFromJSON(jsonValue));
    }

    /**
     * Create SubjectType
     * Create SubjectType
     */
    async createSubjectType(requestParameters: CreateSubjectTypeRequest): Promise<EntSubjectType> {
        const response = await this.createSubjectTypeRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create Teacher
     * Create Teacher
     */
    async createTeacherRaw(requestParameters: CreateTeacherRequest): Promise<runtime.ApiResponse<EntTeacher>> {
        if (requestParameters.teacher === null || requestParameters.teacher === undefined) {
            throw new runtime.RequiredError('teacher','Required parameter requestParameters.teacher was null or undefined when calling createTeacher.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/Teachers`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntCourseToJSON(requestParameters.teacher),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntTeacherFromJSON(jsonValue));
    }

    /**
     * Create Teacher
     * Create Teacher
     */
    async createTeacher(requestParameters: CreateTeacherRequest): Promise<EntTeacher> {
        const response = await this.createTeacherRaw(requestParameters);
        return await response.value();
    }

    /**
     * get courseitem by ID
     * Delete a courseitem entity by ID
     */
    async deleteCourseitemRaw(requestParameters: DeleteCourseitemRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteCourseitem.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/CourseItems/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get courseitem by ID
     * Delete a courseitem entity by ID
     */
    async deleteCourseitem(requestParameters: DeleteCourseitemRequest): Promise<object> {
        const response = await this.deleteCourseitemRaw(requestParameters);
        return await response.value();
    }

    /**
     * get Teacher by ID
     * Delete a Teacher entity by ID
     */
    async deleteTeacherRaw(requestParameters: DeleteTeacherRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteTeacher.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/Teachers/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get Teacher by ID
     * Delete a Teacher entity by ID
     */
    async deleteTeacher(requestParameters: DeleteTeacherRequest): Promise<object> {
        const response = await this.deleteTeacherRaw(requestParameters);
        return await response.value();
    }

    /**
     * get Course by ID
     * Get a Course entity by ID
     */
    async getCourseRaw(requestParameters: GetCourseRequest): Promise<runtime.ApiResponse<EntCourse>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getCourse.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/Courses/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntCourseFromJSON(jsonValue));
    }

    /**
     * get Course by ID
     * Get a Course entity by ID
     */
    async getCourse(requestParameters: GetCourseRequest): Promise<EntCourse> {
        const response = await this.getCourseRaw(requestParameters);
        return await response.value();
    }

    /**
     * get courseitem by ID
     * Get a courseitem entity by ID
     */
    async getCourseitemRaw(requestParameters: GetCourseitemRequest): Promise<runtime.ApiResponse<EntCourseItem>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getCourseitem.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/CourseItems/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntCourseItemFromJSON(jsonValue));
    }

    /**
     * get courseitem by ID
     * Get a courseitem entity by ID
     */
    async getCourseitem(requestParameters: GetCourseitemRequest): Promise<EntCourseItem> {
        const response = await this.getCourseitemRaw(requestParameters);
        return await response.value();
    }

    /**
     * get Subject by ID
     * Get a Subject entity by ID
     */
    async getSubjectRaw(requestParameters: GetSubjectRequest): Promise<runtime.ApiResponse<EntSubject>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getSubject.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/Subjects/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntSubjectFromJSON(jsonValue));
    }

    /**
     * get Subject by ID
     * Get a Subject entity by ID
     */
    async getSubject(requestParameters: GetSubjectRequest): Promise<EntSubject> {
        const response = await this.getSubjectRaw(requestParameters);
        return await response.value();
    }

    /**
     * get SubjectType by ID
     * Get a SubjectType entity by ID
     */
    async getSubjectTypeRaw(requestParameters: GetSubjectTypeRequest): Promise<runtime.ApiResponse<EntSubjectType>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getSubjectType.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/SubjectTypes/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntSubjectTypeFromJSON(jsonValue));
    }

    /**
     * get SubjectType by ID
     * Get a SubjectType entity by ID
     */
    async getSubjectType(requestParameters: GetSubjectTypeRequest): Promise<EntSubjectType> {
        const response = await this.getSubjectTypeRaw(requestParameters);
        return await response.value();
    }

    /**
     * get Teacher by ID
     * Get a Teacher entity by ID
     */
    async getTeacherRaw(requestParameters: GetTeacherRequest): Promise<runtime.ApiResponse<EntTeacher>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getTeacher.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/Teachers/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntTeacherFromJSON(jsonValue));
    }

    /**
     * get Teacher by ID
     * Get a Teacher entity by ID
     */
    async getTeacher(requestParameters: GetTeacherRequest): Promise<EntTeacher> {
        const response = await this.getTeacherRaw(requestParameters);
        return await response.value();
    }

    /**
     * list Course entities
     * List Course entities
     */
    async listCourseRaw(requestParameters: ListCourseRequest): Promise<runtime.ApiResponse<Array<EntCourse>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/Course`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntCourseFromJSON));
    }

    /**
     * list Course entities
     * List Course entities
     */
    async listCourse(requestParameters: ListCourseRequest): Promise<Array<EntCourse>> {
        const response = await this.listCourseRaw(requestParameters);
        return await response.value();
    }

    /**
     * list courseitem entities
     * List courseitem entities
     */
    async listCourseitemRaw(requestParameters: ListCourseitemRequest): Promise<runtime.ApiResponse<Array<EntCourseItem>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/CourseItems`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntCourseItemFromJSON));
    }

    /**
     * list courseitem entities
     * List courseitem entities
     */
    async listCourseitem(requestParameters: ListCourseitemRequest): Promise<Array<EntCourseItem>> {
        const response = await this.listCourseitemRaw(requestParameters);
        return await response.value();
    }

    /**
     * list Subject entities
     * List Subject entities
     */
    async listSubjectRaw(requestParameters: ListSubjectRequest): Promise<runtime.ApiResponse<Array<EntSubject>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/Subjects`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntSubjectFromJSON));
    }

    /**
     * list Subject entities
     * List Subject entities
     */
    async listSubject(requestParameters: ListSubjectRequest): Promise<Array<EntSubject>> {
        const response = await this.listSubjectRaw(requestParameters);
        return await response.value();
    }

    /**
     * list SubjectType entities
     * List SubjectType entities
     */
    async listSubjectTypeRaw(requestParameters: ListSubjectTypeRequest): Promise<runtime.ApiResponse<Array<EntSubjectType>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/SubjectTypes`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntSubjectTypeFromJSON));
    }

    /**
     * list SubjectType entities
     * List SubjectType entities
     */
    async listSubjectType(requestParameters: ListSubjectTypeRequest): Promise<Array<EntSubjectType>> {
        const response = await this.listSubjectTypeRaw(requestParameters);
        return await response.value();
    }

    /**
     * list Teacher entities
     * List Teacher entities
     */
    async listTeacherRaw(requestParameters: ListTeacherRequest): Promise<runtime.ApiResponse<Array<EntTeacher>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/Teachers`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntTeacherFromJSON));
    }

    /**
     * list Teacher entities
     * List Teacher entities
     */
    async listTeacher(requestParameters: ListTeacherRequest): Promise<Array<EntTeacher>> {
        const response = await this.listTeacherRaw(requestParameters);
        return await response.value();
    }

    /**
     * update courseitem by ID
     * Update a courseitem entity by ID
     */
    async updateCourseitemRaw(requestParameters: UpdateCourseitemRequest): Promise<runtime.ApiResponse<EntCourseItem>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateCourseitem.');
        }

        if (requestParameters.courseitem === null || requestParameters.courseitem === undefined) {
            throw new runtime.RequiredError('courseitem','Required parameter requestParameters.courseitem was null or undefined when calling updateCourseitem.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/CourseItems/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: EntCourseItemToJSON(requestParameters.courseitem),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntCourseItemFromJSON(jsonValue));
    }

    /**
     * update courseitem by ID
     * Update a courseitem entity by ID
     */
    async updateCourseitem(requestParameters: UpdateCourseitemRequest): Promise<EntCourseItem> {
        const response = await this.updateCourseitemRaw(requestParameters);
        return await response.value();
    }

    /**
     * update Teacher by ID
     * Update a Teacher entity by ID
     */
    async updateTeacherRaw(requestParameters: UpdateTeacherRequest): Promise<runtime.ApiResponse<EntTeacher>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateTeacher.');
        }

        if (requestParameters.teacher === null || requestParameters.teacher === undefined) {
            throw new runtime.RequiredError('teacher','Required parameter requestParameters.teacher was null or undefined when calling updateTeacher.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/Teachers/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: EntTeacherToJSON(requestParameters.teacher),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntTeacherFromJSON(jsonValue));
    }

    /**
     * update Teacher by ID
     * Update a Teacher entity by ID
     */
    async updateTeacher(requestParameters: UpdateTeacherRequest): Promise<EntTeacher> {
        const response = await this.updateTeacherRaw(requestParameters);
        return await response.value();
    }

}