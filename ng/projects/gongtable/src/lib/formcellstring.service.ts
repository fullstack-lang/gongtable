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

import { FormCellStringDB } from './formcellstring-db';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class FormCellStringService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  FormCellStringServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private formcellstringsUrl: string

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
    this.formcellstringsUrl = origin + '/api/github.com/fullstack-lang/gongtable/go/v1/formcellstrings';
  }

  /** GET formcellstrings from the server */
  getFormCellStrings(GONG__StackPath: string): Observable<FormCellStringDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<FormCellStringDB[]>(this.formcellstringsUrl, { params: params })
      .pipe(
        tap(),
		// tap(_ => this.log('fetched formcellstrings')),
        catchError(this.handleError<FormCellStringDB[]>('getFormCellStrings', []))
      );
  }

  /** GET formcellstring by id. Will 404 if id not found */
  getFormCellString(id: number, GONG__StackPath: string): Observable<FormCellStringDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.formcellstringsUrl}/${id}`;
    return this.http.get<FormCellStringDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched formcellstring id=${id}`)),
      catchError(this.handleError<FormCellStringDB>(`getFormCellString id=${id}`))
    );
  }

  /** POST: add a new formcellstring to the server */
  postFormCellString(formcellstringdb: FormCellStringDB, GONG__StackPath: string): Observable<FormCellStringDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<FormCellStringDB>(this.formcellstringsUrl, formcellstringdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`posted formcellstringdb id=${formcellstringdb.ID}`)
      }),
      catchError(this.handleError<FormCellStringDB>('postFormCellString'))
    );
  }

  /** DELETE: delete the formcellstringdb from the server */
  deleteFormCellString(formcellstringdb: FormCellStringDB | number, GONG__StackPath: string): Observable<FormCellStringDB> {
    const id = typeof formcellstringdb === 'number' ? formcellstringdb : formcellstringdb.ID;
    const url = `${this.formcellstringsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<FormCellStringDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted formcellstringdb id=${id}`)),
      catchError(this.handleError<FormCellStringDB>('deleteFormCellString'))
    );
  }

  /** PUT: update the formcellstringdb on the server */
  updateFormCellString(formcellstringdb: FormCellStringDB, GONG__StackPath: string): Observable<FormCellStringDB> {
    const id = typeof formcellstringdb === 'number' ? formcellstringdb : formcellstringdb.ID;
    const url = `${this.formcellstringsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<FormCellStringDB>(url, formcellstringdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`updated formcellstringdb id=${formcellstringdb.ID}`)
      }),
      catchError(this.handleError<FormCellStringDB>('updateFormCellString'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in FormCellStringService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("FormCellStringService" + error); // log to console instead

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