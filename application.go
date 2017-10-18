package jsonfaas 
import (
	 	//iogos-replace
			"net/http"
			"time"
			"github.com/gorilla/sessions"
			"github.com/gorilla/context"
			"errors"
			"github.com/cheikhshift/db"
			"bytes"
			"encoding/json"
			"fmt"
			"html"
			"html/template"
			"github.com/fatih/color"
			"strings"
			"reflect"
			"unsafe"
			"os"
			"bufio"
			"log"
			"github.com/elazarl/go-bindata-assetfs"
		)
				var store = sessions.NewCookieStore([]byte("a very very very very secret key"))

				type NoStruct struct {
					/* emptystruct */
				}

				func net_sessionGet(key string,s *sessions.Session) string {
					return s.Values[key].(string)
				}

				func UrlAtZ(url,base string) (isURL bool) {
					isURL = strings.Index(url, base) == 0 
					return
				}


				func net_sessionDelete(s *sessions.Session) string {
						//keys := make([]string, len(s.Values))

						//i := 0
						for k := range s.Values {
						   // keys[i] = k.(string)
						    net_sessionRemove(k.(string), s)
						    //i++
						}

					return ""
				}

				func net_sessionRemove(key string,s *sessions.Session) string {
					delete(s.Values, key)
					return ""
				}
				func net_sessionKey(key string,s *sessions.Session) bool {					
				 if _, ok := s.Values[key]; ok {
					    //do something here
				 		return true
					}

					return false
				}

				 

			 func net_add(x,v float64) float64 {
					return v + x				   
			 }

			 func net_subs(x,v float64) float64 {
				   return v - x
			 }

			 func net_multiply(x,v float64) float64 {
				   return v * x
			 }

			 func net_divided(x,v float64) float64 {
				   return v/x
			 }

	

				func net_sessionGetInt(key string,s *sessions.Session) interface{} {
					return s.Values[key]
				}

				func net_sessionSet(key string, value string,s *sessions.Session) string {
					 s.Values[key] = value
					 return ""
				}
				func net_sessionSetInt(key string, value interface{},s *sessions.Session) string {
					 s.Values[key] = value
					 return ""
				}

				func dbDummy() {
					smap := db.O{}
					smap["key"] = "set"
					log.Println(smap)
				}

				
				func net_importcss(s string) string {
					return "<link rel=\"stylesheet\" href=\"" + s + "\" /> "
				}

				func net_importjs(s string) string {
					return "<script type=\"text/javascript\" src=\"" + s + "\" ></script> "
				}



				func formval(s string, r*http.Request) string {
					return r.FormValue(s)
				}


			
				func renderTemplate(w http.ResponseWriter, p *Page)  bool {
				     defer func() {
					        if n := recover(); n != nil {
					           	 color.Red(fmt.Sprintf("Error loading template in path : web%s.tmpl reason : %s", p.R.URL.Path,n)  )
					           	 
					           	 DebugTemplate( w,p.R , fmt.Sprintf("web%s", p.R.URL.Path) )
					           	 w.WriteHeader(http.StatusInternalServerError)
					           	 
						         pag,err := loadPage("/your-500-page" )
						       
						    
						        if err != nil {
						        	log.Println(err.Error())	        	
						        	return
						        }

						         if pag.isResource {
						        	w.Write(pag.Body)
						    	} else {
						    		pag.R = p.R
						         	pag.Session = p.Session
						    		renderTemplate(w, pag) ///your-500-page"
						     
						    	}
					        }
					    }()


				  
				  
				    t := template.New("PageWrapper")
				    t = t.Funcs(template.FuncMap{"a":net_add,"s":net_subs,"m":net_multiply,"d":net_divided,"js" : net_importjs,"css" : net_importcss,"sd" : net_sessionDelete,"sr" : net_sessionRemove,"sc": net_sessionKey,"ss" : net_sessionSet,"sso": net_sessionSetInt,"sgo" : net_sessionGetInt,"sg" : net_sessionGet,"form" : formval,"eq": equalz, "neq" : nequalz, "lte" : netlt,"Testmodel" : net_structTestmodel,"isTestmodel" : net_castTestmodel})
				    t, _ = t.Parse(ReadyTemplate(p.Body))
				    outp := new(bytes.Buffer)
				    err := t.Execute(outp, p)
				    if err != nil {
				        log.Println(err.Error())
				    	DebugTemplate( w,p.R , fmt.Sprintf("web%s", p.R.URL.Path))
				    	w.WriteHeader(http.StatusInternalServerError)
					    w.Header().Set("Content-Type",  "text/html")
						pag,err := loadPage("/your-500-page" )
						 
						 if err != nil {
						        	log.Println(err.Error())	        	
						        	return false
						 }
						 pag.R = p.R
						         pag.Session = p.Session
						         p = nil
						  if pag.isResource {
				        	w.Write(pag.Body)
				    	} else {
				    		renderTemplate(w, pag) // "/your-500-page" 
				     
				    	}
				    	return false
				    } 


				    p.Session.Save(p.R, w)

				    fmt.Fprintf(w, html.UnescapeString(outp.String()) )
				  
				    return true
					
				    
				}

				func MakeHandler(fn func (http.ResponseWriter, *http.Request, string,*sessions.Session)) http.HandlerFunc {
				  return func(w http.ResponseWriter, r *http.Request) {
				  	 
					var session *sessions.Session
				  	var er error
				  	if 	session, er = store.Get(r, "session-"); er != nil {
						session,_ = store.New(r, "session-")
					}
				  	if attmpt := apiAttempt(w,r,session) ;!attmpt {
				      fn(w, r, "",session)
				  	}
				  	
				  	session = nil
				  	context.Clear(r)
				  }
				} 

				func mResponse(v interface{}) string {
					data,_ := json.Marshal(&v)
					return string(data)
				}
				func apiAttempt(w http.ResponseWriter, r *http.Request, session *sessions.Session) (callmet bool) {
					var response string
					response = ""
					
					

					 
				if  isURL := (r.URL.Path == "/test/json" && r.Method == strings.ToUpper("GET") );!callmet && isURL{ 
					
	//Golang code here
		 decoder := json.NewDecoder(r.Body)
		 var t Testmodel
		 err := decoder.Decode(&t)
		 if err != nil {
		    panic(err)
		 } 
		t.FieldOne = "NewValue"
		response = mResponse(t)
	
	
					callmet = true
				}
				

					if callmet {
						session.Save(r,w)
						if response != "" {
							//Unmarshal json
							w.Header().Set("Access-Control-Allow-Origin", "*")
							w.Header().Set("Content-Type",  "application/json")
							w.Write([]byte(response))
						}
						return 
					}
					return
				}
				func SetField(obj interface{}, name string, value interface{}) error {
					structValue := reflect.ValueOf(obj).Elem()
					structFieldValue := structValue.FieldByName(name)

					if !structFieldValue.IsValid() {
						return fmt.Errorf("No such field: %s in obj", name)
					}

					if !structFieldValue.CanSet() {
						return fmt.Errorf("Cannot set %s field value", name)
					}

					structFieldType := structFieldValue.Type()
					val := reflect.ValueOf(value)
					if structFieldType != val.Type() {
						invalidTypeError := errors.New("Provided value type didn't match obj field type")
						return invalidTypeError
					}

					structFieldValue.Set(val)
					return nil
				}
				func DebugTemplate(w http.ResponseWriter,r *http.Request,tmpl string){
					lastline := 0
					linestring := ""
					defer func() {
					       if n := recover(); n != nil {
					           	log.Println()
					           	// log.Println(n)
					           			log.Println("Error on line :", lastline + 1 ,":" + strings.TrimSpace(linestring)) 
					           	 //http.Redirect(w,r,"/your-500-page",307)
					        }
					    }()	

					p,err := loadPage(r.URL.Path)
					filename :=  tmpl  + ".tmpl"
				    body, err := Asset(filename)
				    session, er := store.Get(r, "session-")

				 	if er != nil {
				           session,er = store.New(r,"session-")
				    }
				    p.Session = session
				    p.R = r
				    if err != nil {
				       	log.Print(err)
				    	
				    } else {
				    
				  
				   
				    lines := strings.Split(string(body), "\n")
				   // log.Println( lines )
				    linebuffer := ""
				    waitend := false
				    open := 0
				    for i, line := range lines {
				    	
				    	processd := false
				    	

				    	if strings.Contains(line, "{{with") || strings.Contains(line, "{{ with") || strings.Contains(line, "with}}") || strings.Contains(line, "with }}") || strings.Contains(line, "{{range") || strings.Contains(line, "{{ range") || strings.Contains(line, "range }}") || strings.Contains(line, "range}}") || strings.Contains(line, "{{if") || strings.Contains(line, "{{ if") || strings.Contains(line, "if }}") || strings.Contains(line, "if}}") || strings.Contains(line, "{{block") || strings.Contains(line, "{{ block") || strings.Contains(line, "block }}") || strings.Contains(line, "block}}") {
				    		linebuffer += line
				    		waitend = true
				    		
				    		endstr := ""
				    		processd = true
				    		if !(strings.Contains(line, "{{end") || strings.Contains(line, "{{ end") || strings.Contains(line, "end}}") || strings.Contains(line, "end }}") ) {
				    
				    			open++;
					    		
				    		}
				    		for i := 0; i < open; i++ {
				    			endstr += "\n{{end}}"
				    		}
				    		//exec
				    		outp := new(bytes.Buffer)  
					    	t := template.New("PageWrapper")
					    	t = t.Funcs(template.FuncMap{"a":net_add,"s":net_subs,"m":net_multiply,"d":net_divided,"js" : net_importjs,"css" : net_importcss,"sd" : net_sessionDelete,"sr" : net_sessionRemove,"sc": net_sessionKey,"ss" : net_sessionSet,"sso": net_sessionSetInt,"sgo" : net_sessionGetInt,"sg" : net_sessionGet,"form" : formval,"eq": equalz, "neq" : nequalz, "lte" : netlt,"Testmodel" : net_structTestmodel,"isTestmodel" : net_castTestmodel})
					    	t, _ = t.Parse(ReadyTemplate(body))
					    	lastline = i
					    	linestring =  line
					    	erro := t.Execute(outp, p)
						    if erro != nil {
						   		log.Println("Error on line :", i + 1,line,erro.Error())   
						    } 
				    	}
				   

				    	if waitend && !processd && !( strings.Contains(line, "{{end") || strings.Contains(line, "{{ end") ) {
				    		linebuffer += line

				    		endstr := ""
				    		for i := 0; i < open; i++ {
				    			endstr += "\n{{end}}"
				    		}
				    		//exec
				    		outp := new(bytes.Buffer)  
					    	t := template.New("PageWrapper")
					    	t = t.Funcs(template.FuncMap{"a":net_add,"s":net_subs,"m":net_multiply,"d":net_divided,"js" : net_importjs,"css" : net_importcss,"sd" : net_sessionDelete,"sr" : net_sessionRemove,"sc": net_sessionKey,"ss" : net_sessionSet,"sso": net_sessionSetInt,"sgo" : net_sessionGetInt,"sg" : net_sessionGet,"form" : formval,"eq": equalz, "neq" : nequalz, "lte" : netlt,"Testmodel" : net_structTestmodel,"isTestmodel" : net_castTestmodel})
					    	t, _ = t.Parse(ReadyTemplate(body) )
					    	lastline = i
					    	linestring =  line
					    	erro := t.Execute(outp, p)
						    if erro != nil {
						   		log.Println("Error on line :", i + 1,line,erro.Error())   
						    } 

				    	}



				    	if !waitend && !processd {
				    	outp := new(bytes.Buffer)  
				    	t := template.New("PageWrapper")
				    	t = t.Funcs(template.FuncMap{"a":net_add,"s":net_subs,"m":net_multiply,"d":net_divided,"js" : net_importjs,"css" : net_importcss,"sd" : net_sessionDelete,"sr" : net_sessionRemove,"sc": net_sessionKey,"ss" : net_sessionSet,"sso": net_sessionSetInt,"sgo" : net_sessionGetInt,"sg" : net_sessionGet,"form" : formval,"eq": equalz, "neq" : nequalz, "lte" : netlt,"Testmodel" : net_structTestmodel,"isTestmodel" : net_castTestmodel})
				    	t, _ = t.Parse(ReadyTemplate(body) )
				    	lastline = i
				    	linestring = line
				    	erro := t.Execute(outp, p)
					    if erro != nil {
					   		log.Println("Error on line :", i + 1,line,erro.Error())   
					    }  
						}

						if  !processd && ( strings.Contains(line, "{{end") || strings.Contains(line, "{{ end") ) {
							open--

							if open == 0 {
							waitend = false
				    		
							}
				    	}
				    }
				    
					
				    }

				}

			func DebugTemplatePath(tmpl string, intrf interface{}){
					lastline := 0
					linestring := ""
					defer func() {
					       if n := recover(); n != nil {
					         
					           			log.Println("Error on line :", lastline + 1,":" + strings.TrimSpace(linestring)) 
					           			log.Println(n)
					           	 //http.Redirect(w,r,"/your-500-page",307)
					        }
					    }()	

				
					filename :=  tmpl  
				    body, err := Asset(filename)
				   
				    if err != nil {
				       	log.Print(err)
				    	
				    } else {
				    
				  
				   
				    lines := strings.Split(string(body), "\n")
				   // log.Println( lines )
				    linebuffer := ""
				    waitend := false
				    open := 0
				    for i, line := range lines {
				    	
				    	processd := false

				   		if strings.Contains(line, "{{with") || strings.Contains(line, "{{ with") || strings.Contains(line, "with}}") || strings.Contains(line, "with }}") || strings.Contains(line, "{{range") || strings.Contains(line, "{{ range") || strings.Contains(line, "range }}") || strings.Contains(line, "range}}") || strings.Contains(line, "{{if") || strings.Contains(line, "{{ if") || strings.Contains(line, "if }}") || strings.Contains(line, "if}}") || strings.Contains(line, "{{block") || strings.Contains(line, "{{ block") || strings.Contains(line, "block }}") || strings.Contains(line, "block}}") {
				    		linebuffer += line
				    		waitend = true
				    		
				    		endstr := ""
				    		if !(strings.Contains(line, "{{end") || strings.Contains(line, "{{ end") || strings.Contains(line, "end}}") || strings.Contains(line, "end }}") ) {
				    
				    			open++;
					    		
				    		}

				    		for i := 0; i < open; i++ {
					    			endstr += "\n{{end}}"
					    	}
				    		//exec

				    		processd = true
				    		outp := new(bytes.Buffer)  
					    	t := template.New("PageWrapper")
					    	t = t.Funcs(template.FuncMap{"a":net_add,"s":net_subs,"m":net_multiply,"d":net_divided,"js" : net_importjs,"css" : net_importcss,"sd" : net_sessionDelete,"sr" : net_sessionRemove,"sc": net_sessionKey,"ss" : net_sessionSet,"sso": net_sessionSetInt,"sgo" : net_sessionGetInt,"sg" : net_sessionGet,"form" : formval,"eq": equalz, "neq" : nequalz, "lte" : netlt,"Testmodel" : net_structTestmodel,"isTestmodel" : net_castTestmodel})
					    	t, _ = t.Parse(ReadyTemplate([]byte(fmt.Sprintf("%s%s",linebuffer, endstr))) )
					    	lastline = i
					    	linestring =  line	    	
					    	erro := t.Execute(outp, intrf)
						    if erro != nil {
						   		log.Println("Error on line :", i + 1,line,erro.Error())   
						    } 
				    	}



				    	if waitend && !processd && !(strings.Contains(line, "{{end") || strings.Contains(line, "{{ end") || strings.Contains(line, "end}}") || strings.Contains(line, "end }}") )  {
				    		linebuffer += line

				    		endstr := ""
				    		for i := 0; i < open; i++ {
				    			endstr += "\n{{end}}"
				    		}
				    		//exec
				    		outp := new(bytes.Buffer)  
					    	t := template.New("PageWrapper")
					    	t = t.Funcs(template.FuncMap{"a":net_add,"s":net_subs,"m":net_multiply,"d":net_divided,"js" : net_importjs,"css" : net_importcss,"sd" : net_sessionDelete,"sr" : net_sessionRemove,"sc": net_sessionKey,"ss" : net_sessionSet,"sso": net_sessionSetInt,"sgo" : net_sessionGetInt,"sg" : net_sessionGet,"form" : formval,"eq": equalz, "neq" : nequalz, "lte" : netlt,"Testmodel" : net_structTestmodel,"isTestmodel" : net_castTestmodel})
					    	t, _ = t.Parse(ReadyTemplate([]byte(fmt.Sprintf("%s%s",linebuffer, endstr))) )
					    	lastline = i
					    	linestring =  line
					    	erro := t.Execute(outp, intrf)
						    if erro != nil {
						   		log.Println("Error on line :", i + 1,line,erro.Error())   
						    } 

				    	}



				    	if !waitend && !processd {
				    	outp := new(bytes.Buffer)  
				    	t := template.New("PageWrapper")
				    	t = t.Funcs(template.FuncMap{"a":net_add,"s":net_subs,"m":net_multiply,"d":net_divided,"js" : net_importjs,"css" : net_importcss,"sd" : net_sessionDelete,"sr" : net_sessionRemove,"sc": net_sessionKey,"ss" : net_sessionSet,"sso": net_sessionSetInt,"sgo" : net_sessionGetInt,"sg" : net_sessionGet,"form" : formval,"eq": equalz, "neq" : nequalz, "lte" : netlt,"Testmodel" : net_structTestmodel,"isTestmodel" : net_castTestmodel})
					    t, _ = t.Parse(ReadyTemplate([]byte(fmt.Sprintf("%s%s",linebuffer))) )
				    	lastline = i
				    	linestring = line
				    	erro := t.Execute(outp, intrf)
					    if erro != nil {
					   		log.Println("Error on line :", i + 1,line,erro.Error())   
					    }  
						}

						if  !processd && (strings.Contains(line, "{{end") || strings.Contains(line, "{{ end") || strings.Contains(line, "end}}") || strings.Contains(line, "end }}") ) {
							open--

							if open == 0 {
							waitend = false
				    		
							}
				    	}
				    }
				    
					
				    }

				}
			func Handler(w http.ResponseWriter, r *http.Request, contxt string,session *sessions.Session) {
				  var p *Page
				  p,err := loadPage(r.URL.Path)
				  
				  if err != nil {	
				  		log.Println(err.Error())
				  		
				        w.WriteHeader(http.StatusNotFound)				  	
				       	
				        pag,err := loadPage("/your-404-page")
				        
				        if err != nil {
				        	log.Println(err.Error())
				        	//context.Clear(r)
				        	return
				        }
				         pag.R = r
						         pag.Session = session
						         p = nil
				        if pag.isResource {
				        	w.Write(pag.Body)
				    	} else {
				    		renderTemplate(w, pag) //"/your-500-page" 
				    	}
				        return
				  }

				   
				  if !p.isResource {
				  		w.Header().Set("Content-Type",  "text/html")
				  		p.Session = session
				  		p.R = r
				      	renderTemplate(w, p) //fmt.Sprintf("web%s", r.URL.Path)
				     
				     // log.Println(w)
				  } else {
				  		w.Header().Set("Cache-Control",  "public")
				  		if strings.Contains(r.URL.Path, ".css") {
				  	  		w.Header().Add("Content-Type",  "text/css")
				  	  	} else if strings.Contains(r.URL.Path, ".js") {
				  	  		w.Header().Add("Content-Type",  "application/javascript")
				  	  	} else {
				  	  	w.Header().Add("Content-Type",  http.DetectContentType(p.Body))
				  	  	}
				  	 
				  	 
				      w.Write(p.Body)
				  }

				  p.R = nil
				  p.Session = nil
				  p = nil
				  //context.Clear(r)
				  
				}

				func loadPage(title string) (*Page,error) {
				   
				    if roottitle := (title == "/"); roottitle  {
				    	webbase := "web/"
					    	fname := fmt.Sprintf("%s%s", webbase, "index.html")
					    	body, err := Asset(fname)
					    	if err != nil {
					    		fname = fmt.Sprintf("%s%s", webbase, "index.tmpl")
					    		body , err = Asset(fname)
					    		if err != nil {
					    			return nil,err
					    		}
					    		return  &Page{ Body: body,isResource: false}, nil
					    	}

					    	return  &Page{ Body: body,isResource: true}, nil
					    		    		
				     } 
				     
				    filename := fmt.Sprintf("web%s.tmpl", title)

				   if body, err := Asset(filename) ;err != nil {
				    	 filename = fmt.Sprintf("web%s.html", title) 
				    	
				    	if  body, err = Asset(filename); err != nil {
				         filename = fmt.Sprintf("web%s", title) 
				         
				         if  body, err = Asset(filename); err != nil {
				            return nil, err
				         } else {
				          if strings.Contains(title, ".tmpl")  {
				              return nil,nil
				          }
				          return &Page{ Body: body,isResource: true }, nil
				         }
				      } else {
				         return &Page{ Body: body,isResource: true}, nil
				      }
				    } else {
				    	  return &Page{Body: body,isResource:false}, nil
				    }
 
				       
				  
				 } 
				
				   
				
				func BytesToString(b []byte) string {
				    bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
				    sh := reflect.StringHeader{bh.Data, bh.Len}
				    return *(*string)(unsafe.Pointer(&sh))
				}
				func equalz(args ...interface{}) bool {
		    	    if args[0] == args[1] {
		        	return true;
				    }
				    return false;
				 }
				 func nequalz(args ...interface{}) bool {
				    if args[0] != args[1] {
				        return true;
				    }
				    return false;
				 }

				 func netlt(x,v float64) bool {
				    if x < v {
				        return true;
				    }
				    return false;
				 }
				 func netgt(x,v float64) bool {
				    if x > v {
				        return true;
				    }
				    return false;
				 }
				 func netlte(x,v float64) bool {
				    if x <= v {
				        return true;
				    }
				    return false;
				 }

				 func GetLine(fname string , match string )  int {
					intx := 0
					file, err := os.Open(fname)
								if err != nil {
									color.Red("Could not find a source file")
																		           		return -1
												    }
								defer file.Close()

								scanner := bufio.NewScanner(file)
								for scanner.Scan() {
									intx = intx + 1
									if strings.Contains(scanner.Text(), match ) {
												    		
												    		return intx
												    	}

								}


					return -1
				}
				 func netgte(x,v float64) bool {
				    if x >= v {
				        return true;
				    }
				    return false;
				 }
				 type Page struct {
					    Title string
					    Body  []byte
					    request *http.Request
					    isResource bool
					    R *http.Request
					    Session *sessions.Session
				 }

				 func ReadyTemplate(body []byte) string { return strings.Replace(strings.Replace(strings.Replace(string(body), "/{", "\"{",-1),"}/", "}\"",-1 ) ,"`", "\"" ,-1) }
				 
			type Testmodel struct {
	//interface fields here
	FieldOne string
	FieldTwo int
	FieldThree []string
	
			}

			func  net_castTestmodel(args ...interface{}) *Testmodel  {
				
				s := Testmodel{}
				mapp := args[0].(db.O)
				if _, ok := mapp["_id"]; ok {
					mapp["Id"] = mapp["_id"]
				}
				data,_ := json.Marshal(&mapp)
				
				err := json.Unmarshal(data, &s) 
				if err != nil {
					log.Println(err.Error())
				}
				
				return &s
			}
			func net_structTestmodel() *Testmodel{ return &Testmodel{} }
			func dummy_timer(){
				dg := time.Second *5
				log.Println(dg)
			}
		
		func MakeFSHandle(root string){
			http.Handle("/dist/",  http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: root}))
		}
		