// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { FormFieldDateDB } from './formfielddate-db';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class FormFieldDateService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  FormFieldDateServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private formfielddatesUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.formfielddatesUrl = origin + '/api/github.com/fullstack-lang/gongtable/go/v1/formfielddates';
  }

  /** GET formfielddates from the server */
  getFormFieldDates(GONG__StackPath: string): Observable<FormFieldDateDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<FormFieldDateDB[]>(this.formfielddatesUrl, { params: params })
      .pipe(
        tap(),
		// tap(_ => this.log('fetched formfielddates')),
        catchError(this.handleError<FormFieldDateDB[]>('getFormFieldDates', []))
      );
  }

  /** GET formfielddate by id. Will 404 if id not found */
  getFormFieldDate(id: number, GONG__StackPath: string): Observable<FormFieldDateDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.formfielddatesUrl}/${id}`;
    return this.http.get<FormFieldDateDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched formfielddate id=${id}`)),
      catchError(this.handleError<FormFieldDateDB>(`getFormFieldDate id=${id}`))
    );
  }

  /** POST: add a new formfielddate to the server */
  postFormFieldDate(formfielddatedb: FormFieldDateDB, GONG__StackPath: string): Observable<FormFieldDateDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<FormFieldDateDB>(this.formfielddatesUrl, formfielddatedb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`posted formfielddatedb id=${formfielddatedb.ID}`)
      }),
      catchError(this.handleError<FormFieldDateDB>('postFormFieldDate'))
    );
  }

  /** DELETE: delete the formfielddatedb from the server */
  deleteFormFieldDate(formfielddatedb: FormFieldDateDB | number, GONG__StackPath: string): Observable<FormFieldDateDB> {
    const id = typeof formfielddatedb === 'number' ? formfielddatedb : formfielddatedb.ID;
    const url = `${this.formfielddatesUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<FormFieldDateDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted formfielddatedb id=${id}`)),
      catchError(this.handleError<FormFieldDateDB>('deleteFormFieldDate'))
    );
  }

  /** PUT: update the formfielddatedb on the server */
  updateFormFieldDate(formfielddatedb: FormFieldDateDB, GONG__StackPath: string): Observable<FormFieldDateDB> {
    const id = typeof formfielddatedb === 'number' ? formfielddatedb : formfielddatedb.ID;
    const url = `${this.formfielddatesUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<FormFieldDateDB>(url, formfielddatedb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`updated formfielddatedb id=${formfielddatedb.ID}`)
      }),
      catchError(this.handleError<FormFieldDateDB>('updateFormFieldDate'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in FormFieldDateService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("FormFieldDateService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
      console.log(message)
  }
}