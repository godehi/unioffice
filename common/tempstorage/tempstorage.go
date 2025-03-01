//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package tempstorage ;import _d "io";

// Add reads a file from a disk and adds it to the storage.
func Add (path string )error {return _e .Add (path )};type storage interface{Open (_g string )(File ,error );TempFile (_a ,_dd string )(File ,error );TempDir (_ge string )(string ,error );RemoveAll (_c string )error ;Add (_dc string )error ;};

// TempDir creates a name for a new temp directory using a pattern argument.
func TempDir (pattern string )(string ,error ){return _e .TempDir (pattern )};

// Open returns tempstorage File object by name.
func Open (path string )(File ,error ){return _e .Open (path )};

// TempFile creates new empty file in the storage and returns it.
func TempFile (dir ,pattern string )(File ,error ){return _e .TempFile (dir ,pattern )};var _e storage ;

// File is a representation of a storage file
// with Read, Write, Close and Name methods identical to os.File.
type File interface{_d .Reader ;_d .ReaderAt ;_d .Writer ;_d .Closer ;Name ()string ;};

// SetAsStorage changes temporary storage to newStorage.
func SetAsStorage (newStorage storage ){_e =newStorage };

// RemoveAll removes all files according to the dir argument prefix.
func RemoveAll (dir string )error {return _e .RemoveAll (dir )};