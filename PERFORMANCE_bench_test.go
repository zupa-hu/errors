
package errors

import (
	"testing"
)

func BenchmarkOk_return(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = func() (Error) {
			return nil
		}()
	}

	_ = Err
}
func BenchmarkOk_catch(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = Catch(func() {
			// NOP
		})
	}
	
	_ = Err
}
func BenchmarkOk_context(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = InContext(func(e *Context) {
			// NOP
		})
	}
	
	_ = Err
}




func BenchmarkErr_return(b *testing.B) {
	Err := Type("ErrTest").ServerError("woof!")
	var Err2 Error

	for i:=0; i<b.N; i++ {
		Err2 = func() (Error) {
			return Err
		}()
	}
	
	_ = Err2
}
func BenchmarkErr_catch(b *testing.B) {
	Err := Type("ErrTest").ServerError("woof!")
	var Err2 Error

	for i:=0; i<b.N; i++ {
		Err2 = Catch(func() {
			panic(Err)
		})
	}
	
	_ = Err2
}
func BenchmarkErr_context(b *testing.B) {
	Err := Type("ErrTest").ServerError("woof!")
	var Err2 Error

	for i:=0; i<b.N; i++ {
		Err2 = InContext(func(e *Context) {
			e.Throw(Err)
		})
	}
	
	_ = Err2
}



func BenchmarkDeep_return(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = func() (Error) {
			return func() (Error) {
				return func() (Error) {
					return func() (Error) {
						return func() (Error) {
							return func() (Error) {
								return func() (Error) {
									return func() (Error) {
										return func() (Error) {
											return func() (Error) {
												return nil
											}()
										}()
									}()
								}()
							}()
						}()
					}()
				}()
			}()
		}()
	}

	_ = Err
}
func BenchmarkDeep_catch(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = Catch(func() {
			func() {
				func() {
					func() {
						func() {
							func() {
								func() {
									func() {
										func() {
											func() {
												// NOP
											}()
										}()
									}()
								}()
							}()
						}()
					}()
				}()
			}()
		})
	}
	
	_ = Err
}
func BenchmarkDeep_context(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = InContext(func(e *Context) {
			func() {
				func() {
					func() {
						func() {
							func() {
								func() {
									func() {
										func() {
											func() {
												// NOP
											}()
										}()
									}()
								}()
							}()
						}()
					}()
				}()
			}()
		})
	}
	
	_ = Err
}


func BenchmarkUseFulMany_10_return(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = func() (Error) {
			for j:=0; j<10; j++ {
				bar, Err := func(n int) (string, Error) {
					if n == 99999999 {
						return "", Clientf("unexpected n [%v]", n)
					}
					return "foo", nil
				}(j)
				if Err != nil { return Err }
				_ = bar
			}

			return nil
		}()
	}

	_ = Err
}
func BenchmarkUseFulMany_10_catch(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = Catch(func() {
			for j:=0; j<10; j++ {
				bar := func(n int) (string) {
					if n == 99999999 {
						panic(Clientf("unexpected n [%v]", n))
					}
					return "foo"
				}(j)
				_ = bar
			}
		})
	}
	
	_ = Err
}
func BenchmarkUseFulMany_10_context(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = InContext(func(e *Context) {
			for j:=0; j<10; j++ {
				bar := func(n int) (string) {
					if n == 99999999 {
						e.Clientf("unexpected n [%v]", n)
					}
					return "foo"
				}(j)
				_ = bar
			}
		})
	}
	
	_ = Err
}

func BenchmarkUseFulMany_100_return(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = func() (Error) {
			for j:=0; j<100; j++ {
				bar, Err := func(n int) (string, Error) {
					if n == 99999999 {
						return "", Clientf("unexpected n [%v]", n)
					}
					return "foo", nil
				}(j)
				if Err != nil { return Err }
				_ = bar
			}

			return nil
		}()
	}

	_ = Err
}
func BenchmarkUseFulMany_100_catch(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = Catch(func() {
			for j:=0; j<100; j++ {
				bar := func(n int) (string) {
					if n == 99999999 {
						panic(Clientf("unexpected n [%v]", n))
					}
					return "foo"
				}(j)
				_ = bar
			}
		})
	}
	
	_ = Err
}
func BenchmarkUseFulMany_100_context(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = InContext(func(e *Context) {
			for j:=0; j<100; j++ {
				bar := func(n int) (string) {
					if n == 99999999 {
						e.Clientf("unexpected n [%v]", n)
					}
					return "foo"
				}(j)
				_ = bar
			}
		})
	}
	
	_ = Err
}

func BenchmarkUseFulMany_1000_return(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = func() (Error) {
			for j:=0; j<1000; j++ {
				bar, Err := func(n int) (string, Error) {
					if n == 99999999 {
						return "", Clientf("unexpected n [%v]", n)
					}
					return "foo", nil
				}(j)
				if Err != nil { return Err }
				_ = bar
			}

			return nil
		}()
	}

	_ = Err
}
func BenchmarkUseFulMany_1000_catch(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = Catch(func() {
			for j:=0; j<1000; j++ {
				bar := func(n int) (string) {
					if n == 99999999 {
						panic(Clientf("unexpected n [%v]", n))
					}
					return "foo"
				}(j)
				_ = bar
			}
		})
	}
	
	_ = Err
}
func BenchmarkUseFulMany_1000_context(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = InContext(func(e *Context) {
			for j:=0; j<1000; j++ {
				bar := func(n int) (string) {
					if n == 99999999 {
						e.Clientf("unexpected n [%v]", n)
					}
					return "foo"
				}(j)
				_ = bar
			}
		})
	}
	
	_ = Err
}

func BenchmarkUseFulMany_10000_return(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = func() (Error) {
			for j:=0; j<10000; j++ {
				bar, Err := func(n int) (string, Error) {
					if n == 99999999 {
						return "", Clientf("unexpected n [%v]", n)
					}
					return "foo", nil
				}(j)
				if Err != nil { return Err }
				_ = bar
			}

			return nil
		}()
	}

	_ = Err
}
func BenchmarkUseFulMany_10000_catch(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = Catch(func() {
			for j:=0; j<10000; j++ {
				bar := func(n int) (string) {
					if n == 99999999 {
						panic(Clientf("unexpected n [%v]", n))
					}
					return "foo"
				}(j)
				_ = bar
			}
		})
	}
	
	_ = Err
}
func BenchmarkUseFulMany_10000_context(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = InContext(func(e *Context) {
			for j:=0; j<10000; j++ {
				bar := func(n int) (string) {
					if n == 99999999 {
						e.Clientf("unexpected n [%v]", n)
					}
					return "foo"
				}(j)
				_ = bar
			}
		})
	}
	
	_ = Err
}


