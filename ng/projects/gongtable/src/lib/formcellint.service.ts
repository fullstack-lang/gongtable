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

import { FormCellIntDB } from './formcellint-db';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class FormCellIntService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  FormCellIntServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private formcellintsUrl: string

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
    this.formcellintsUrl = origin + '/api/github.com/fullstack-lang/gongtable/go/v1/formcellints';
  }

  /** GET formcellints from the server */
  getFormCellInts(GONG__StackPath: string): Observable<FormCellIntDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<FormCellIntDB[]>(this.formcellintsUrl, { params: params })
      .pipe(
        tap(),
		// tap(_ => this.log('fetched formcellints')),
        catchError(this.handleError<FormCellIntDB[]>('getFormCellInts', []))
      );
  }

  /** GET formcellint by id. Will 404 if id not found */
  getFormCellInt(id: number, GONG__StackPath: string): Observable<FormCellIntDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.formcellintsUrl}/${id}`;
    return this.http.get<FormCellIntDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched formcellint id=${id}`)),
      catchError(this.handleError<FormCellIntDB>(`getFormCellInt id=${id}`))
    );
  }

  /** POST: add a new formcellint to the server */
  postFormCellInt(formcellintdb: FormCellIntDB, GONG__StackPath: string): Observable<FormCellIntDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<FormCellIntDB>(this.formcellintsUrl, formcellintdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`posted formcellintdb id=${formcellintdb.ID}`)
      }),
      catchError(this.handleError<FormCellIntDB>('postFormCellInt'))
    );
  }

  /** DELETE: delete the formcellintdb from the server */
  deleteFormCellInt(formcellintdb: FormCellIntDB | number, GONG__StackPath: string): Observable<FormCellIntDB> {
    const id = typeof formcellintdb === 'number' ? formcellintdb : formcellintdb.ID;
    const url = `${this.formcellintsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<FormCellIntDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted formcellintdb id=${id}`)),
      catchError(this.handleError<FormCellIntDB>('deleteFormCellInt'))
    );
  }

  /** PUT: update the formcellintdb on the server */
  updateFormCellInt(formcellintdb: FormCellIntDB, GONG__StackPath: string): Observable<FormCellIntDB> {
    const id = typeof formcellintdb === 'number' ? formcellintdb : formcellintdb.ID;
    const url = `${this.formcellintsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<FormCellIntDB>(url, formcellintdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`updated formcellintdb id=${formcellintdb.ID}`)
      }),
      catchError(this.handleError<FormCellIntDB>('updateFormCellInt'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in FormCellIntService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("FormCellIntService" + error); // log to console instead

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