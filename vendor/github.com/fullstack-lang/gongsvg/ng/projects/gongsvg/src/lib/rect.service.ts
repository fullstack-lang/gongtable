// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { RectDB } from './rect-db'
import { Rect, CopyRectToRectDB } from './rect'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { AnimateDB } from './animate-db'
import { RectAnchoredTextDB } from './rectanchoredtext-db'
import { RectAnchoredRectDB } from './rectanchoredrect-db'
import { RectAnchoredPathDB } from './rectanchoredpath-db'

@Injectable({
  providedIn: 'root'
})
export class RectService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  RectServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private rectsUrl: string

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
    this.rectsUrl = origin + '/api/github.com/fullstack-lang/gongsvg/go/v1/rects';
  }

  /** GET rects from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectDB[]> {
    return this.getRects(GONG__StackPath, frontRepo)
  }
  getRects(GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<RectDB[]>(this.rectsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<RectDB[]>('getRects', []))
      );
  }

  /** GET rect by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectDB> {
    return this.getRect(id, GONG__StackPath, frontRepo)
  }
  getRect(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.rectsUrl}/${id}`;
    return this.http.get<RectDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched rect id=${id}`)),
      catchError(this.handleError<RectDB>(`getRect id=${id}`))
    );
  }

  // postFront copy rect to a version with encoded pointers and post to the back
  postFront(rect: Rect, GONG__StackPath: string): Observable<RectDB> {
    let rectDB = new RectDB
    CopyRectToRectDB(rect, rectDB)
    const id = typeof rectDB === 'number' ? rectDB : rectDB.ID
    const url = `${this.rectsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<RectDB>(url, rectDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<RectDB>('postRect'))
    );
  }
  
  /** POST: add a new rect to the server */
  post(rectdb: RectDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectDB> {
    return this.postRect(rectdb, GONG__StackPath, frontRepo)
  }
  postRect(rectdb: RectDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<RectDB>(this.rectsUrl, rectdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted rectdb id=${rectdb.ID}`)
      }),
      catchError(this.handleError<RectDB>('postRect'))
    );
  }

  /** DELETE: delete the rectdb from the server */
  delete(rectdb: RectDB | number, GONG__StackPath: string): Observable<RectDB> {
    return this.deleteRect(rectdb, GONG__StackPath)
  }
  deleteRect(rectdb: RectDB | number, GONG__StackPath: string): Observable<RectDB> {
    const id = typeof rectdb === 'number' ? rectdb : rectdb.ID;
    const url = `${this.rectsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<RectDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted rectdb id=${id}`)),
      catchError(this.handleError<RectDB>('deleteRect'))
    );
  }

  // updateFront copy rect to a version with encoded pointers and update to the back
  updateFront(rect: Rect, GONG__StackPath: string): Observable<RectDB> {
    let rectDB = new RectDB
    CopyRectToRectDB(rect, rectDB)
    const id = typeof rectDB === 'number' ? rectDB : rectDB.ID
    const url = `${this.rectsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<RectDB>(url, rectDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<RectDB>('updateRect'))
    );
  }

  /** PUT: update the rectdb on the server */
  update(rectdb: RectDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectDB> {
    return this.updateRect(rectdb, GONG__StackPath, frontRepo)
  }
  updateRect(rectdb: RectDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectDB> {
    const id = typeof rectdb === 'number' ? rectdb : rectdb.ID;
    const url = `${this.rectsUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<RectDB>(url, rectdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated rectdb id=${rectdb.ID}`)
      }),
      catchError(this.handleError<RectDB>('updateRect'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in RectService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("RectService" + error); // log to console instead

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
