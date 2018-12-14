
package errors

import (
	"testing"
)

func BenchmarkCatch_okReturn(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = func() (Error) {
			return nil
		}()
	}

	_ = Err
}
func BenchmarkCatch_okCatch(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = Catch(func() {
			// NOP
		})
	}
	
	_ = Err
}



func BenchmarkCatch_errReturn(b *testing.B) {
	Err := Type("ErrTest").ServerError("woof!")
	var Err2 Error

	for i:=0; i<b.N; i++ {
		Err2 = func() (Error) {
			return Err
		}()
	}
	
	_ = Err2
}
func BenchmarkCatch_errCatch(b *testing.B) {
	Err := Type("ErrTest").ServerError("woof!")
	var Err2 Error

	for i:=0; i<b.N; i++ {
		Err2 = Catch(func() {
			panic(Err)
		})
	}
	
	_ = Err2
}



func BenchmarkCatch_deepReturn(b *testing.B) {
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
func BenchmarkCatch_deepCatch(b *testing.B) {
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



func BenchmarkCatch_many_1000_Return(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = func() (Error) {
			for j:=0; j<1000; j++ {
				Err := func() (Error) {
					return nil
				}()
				if Err != nil { return Err }
			}

			return nil
		}()
	}

	_ = Err
}
func BenchmarkCatch_many_1000_Catch(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = Catch(func() {
			for j:=0; j<1000; j++ {
				func() {
					// NOP
				}()
			}
		})
	}
	
	_ = Err
}
func BenchmarkCatch_many_100_Return(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = func() (Error) {
			for j:=0; j<100; j++ {
				Err := func() (Error) {
					return nil
				}()
				if Err != nil { return Err }
			}

			return nil
		}()
	}

	_ = Err
}
func BenchmarkCatch_many_100_Catch(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = Catch(func() {
			for j:=0; j<100; j++ {
				func() {
					// NOP
				}()
			}
		})
	}
	
	_ = Err
}


func BenchmarkCatch_many_10_Return(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = func() (Error) {
			for j:=0; j<10; j++ {
				Err := func() (Error) {
					return nil
				}()
				if Err != nil { return Err }
			}

			return nil
		}()
	}

	_ = Err
}
func BenchmarkCatch_many_10_Catch(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = Catch(func() {
			for j:=0; j<10; j++ {
				func() {
					// NOP
				}()
			}
		})
	}
	
	_ = Err
}


func BenchmarkCatch_useFulMany_100_Return(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = func() (Error) {
			for j:=0; j<100; j++ {
				bar, Err := func() (string, Error) {
					return "foo", nil
				}()
				if Err != nil { return Err }
				_ = bar
			}

			return nil
		}()
	}

	_ = Err
}
func BenchmarkCatch_useFulMany_100_Catch(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = Catch(func() {
			for j:=0; j<100; j++ {
				bar := func() (string) {
					return "foo"
				}()
				_ = bar
			}
		})
	}
	
	_ = Err
}

func BenchmarkCatch_useFulMany_1000_Return(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = func() (Error) {
			for j:=0; j<1000; j++ {
				bar, Err := func() (string, Error) {
					return "foo", nil
				}()
				if Err != nil { return Err }
				_ = bar
			}

			return nil
		}()
	}

	_ = Err
}
func BenchmarkCatch_useFulMany_1000_Catch(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = Catch(func() {
			for j:=0; j<1000; j++ {
				bar := func() (string) {
					return "foo"
				}()
				_ = bar
			}
		})
	}
	
	_ = Err
}

func BenchmarkCatch_useFulMany_10000_Return(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = func() (Error) {
			for j:=0; j<10000; j++ {
				bar, Err := func() (string, Error) {
					return "foo", nil
				}()
				if Err != nil { return Err }
				_ = bar
			}

			return nil
		}()
	}

	_ = Err
}
func BenchmarkCatch_useFulMany_10000_Catch(b *testing.B) {
	var Err Error

	for i:=0; i<b.N; i++ {
		Err = Catch(func() {
			for j:=0; j<10000; j++ {
				bar := func() (string) {
					return "foo"
				}()
				_ = bar
			}
		})
	}
	
	_ = Err
}