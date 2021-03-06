// +build !minimal

package qtnetwork

// /usr/include/qt/QtNetwork/qdnslookup.h
// #include <qdnslookup.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QDnsLookup struct {
	*qtcore.QObject
}
type QDnsLookup_ITF interface {
	qtcore.QObject_ITF
	QDnsLookup_PTR() *QDnsLookup
}

func (ptr *QDnsLookup) QDnsLookup_PTR() *QDnsLookup { return ptr }

func (this *QDnsLookup) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QDnsLookup) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQDnsLookupFromPointer(cthis unsafe.Pointer) *QDnsLookup {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QDnsLookup{bcthis0}
}
func (*QDnsLookup) NewFromPointer(cthis unsafe.Pointer) *QDnsLookup {
	return NewQDnsLookupFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:188
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QDnsLookup) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDnsLookup10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qdnslookup.h:223
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDnsLookup(QObject *)

/*
Constructs a QDnsLookup object and sets parent as the parent object.

The type property will default to QDnsLookup::A.
*/
func (*QDnsLookup) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QDnsLookup {
	return NewQDnsLookup(parent)
}
func NewQDnsLookup(parent qtcore.QObject_ITF /*777 QObject **/) *QDnsLookup {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDnsLookupC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDnsLookupFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDnsLookup")
	return gothis
}

// /usr/include/qt/QtNetwork/qdnslookup.h:223
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDnsLookup(QObject *)

/*
Constructs a QDnsLookup object and sets parent as the parent object.

The type property will default to QDnsLookup::A.
*/
func (*QDnsLookup) NewForInheritp() *QDnsLookup {
	return NewQDnsLookupp()
}
func NewQDnsLookupp() *QDnsLookup {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDnsLookupC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDnsLookupFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDnsLookup")
	return gothis
}

// /usr/include/qt/QtNetwork/qdnslookup.h:224
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDnsLookup(QDnsLookup::Type, const QString &, QObject *)

/*
Constructs a QDnsLookup object and sets parent as the parent object.

The type property will default to QDnsLookup::A.
*/
func (*QDnsLookup) NewForInherit1(type_ int, name string, parent qtcore.QObject_ITF /*777 QObject **/) *QDnsLookup {
	return NewQDnsLookup1(type_, name, parent)
}
func NewQDnsLookup1(type_ int, name string, parent qtcore.QObject_ITF /*777 QObject **/) *QDnsLookup {
	var tmpArg1 = qtcore.NewQString5(name)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg2 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDnsLookupC2ENS_4TypeERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, type_, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDnsLookupFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDnsLookup")
	return gothis
}

// /usr/include/qt/QtNetwork/qdnslookup.h:224
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDnsLookup(QDnsLookup::Type, const QString &, QObject *)

/*
Constructs a QDnsLookup object and sets parent as the parent object.

The type property will default to QDnsLookup::A.
*/
func (*QDnsLookup) NewForInherit1p(type_ int, name string) *QDnsLookup {
	return NewQDnsLookup1p(type_, name)
}
func NewQDnsLookup1p(type_ int, name string) *QDnsLookup {
	var tmpArg1 = qtcore.NewQString5(name)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, QObject *=Pointer, QObject=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDnsLookupC2ENS_4TypeERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, type_, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDnsLookupFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDnsLookup")
	return gothis
}

// /usr/include/qt/QtNetwork/qdnslookup.h:225
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QDnsLookup(QDnsLookup::Type, const QString &, const QHostAddress &, QObject *)

/*
Constructs a QDnsLookup object and sets parent as the parent object.

The type property will default to QDnsLookup::A.
*/
func (*QDnsLookup) NewForInherit2(type_ int, name string, nameserver QHostAddress_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QDnsLookup {
	return NewQDnsLookup2(type_, name, nameserver, parent)
}
func NewQDnsLookup2(type_ int, name string, nameserver QHostAddress_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QDnsLookup {
	var tmpArg1 = qtcore.NewQString5(name)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if nameserver != nil && nameserver.QHostAddress_PTR() != nil {
		convArg2 = nameserver.QHostAddress_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg3 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDnsLookupC2ENS_4TypeERK7QStringRK12QHostAddressP7QObject", qtrt.FFI_TYPE_POINTER, type_, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDnsLookupFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDnsLookup")
	return gothis
}

// /usr/include/qt/QtNetwork/qdnslookup.h:225
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QDnsLookup(QDnsLookup::Type, const QString &, const QHostAddress &, QObject *)

/*
Constructs a QDnsLookup object and sets parent as the parent object.

The type property will default to QDnsLookup::A.
*/
func (*QDnsLookup) NewForInherit2p(type_ int, name string, nameserver QHostAddress_ITF) *QDnsLookup {
	return NewQDnsLookup2p(type_, name, nameserver)
}
func NewQDnsLookup2p(type_ int, name string, nameserver QHostAddress_ITF) *QDnsLookup {
	var tmpArg1 = qtcore.NewQString5(name)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if nameserver != nil && nameserver.QHostAddress_PTR() != nil {
		convArg2 = nameserver.QHostAddress_PTR().GetCthis()
	}
	// arg: 3, QObject *=Pointer, QObject=Record, , Invalid
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDnsLookupC2ENS_4TypeERK7QStringRK12QHostAddressP7QObject", qtrt.FFI_TYPE_POINTER, type_, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDnsLookupFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDnsLookup")
	return gothis
}

// /usr/include/qt/QtNetwork/qdnslookup.h:226
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDnsLookup()

/*

 */
func DeleteQDnsLookup(this *QDnsLookup) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDnsLookupD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:228
// index:0
// Public Visibility=Default Availability=Available
// [4] QDnsLookup::Error error() const

/*

 */
func (this *QDnsLookup) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDnsLookup5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:229
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*

 */
func (this *QDnsLookup) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDnsLookup11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qdnslookup.h:230
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFinished() const

/*
Returns whether the reply has finished or was aborted.
*/
func (this *QDnsLookup) IsFinished() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDnsLookup10isFinishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qdnslookup.h:232
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name() const

/*

 */
func (this *QDnsLookup) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDnsLookup4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qdnslookup.h:233
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setName(const QString &)

/*

 */
func (this *QDnsLookup) SetName(name string) {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDnsLookup7setNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:235
// index:0
// Public Visibility=Default Availability=Available
// [4] QDnsLookup::Type type() const

/*

 */
func (this *QDnsLookup) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDnsLookup4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:236
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setType(QDnsLookup::Type)

/*

 */
func (this *QDnsLookup) SetType(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDnsLookup7setTypeENS_4TypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:238
// index:0
// Public Visibility=Default Availability=Available
// [8] QHostAddress nameserver() const

/*

 */
func (this *QDnsLookup) Nameserver() *QHostAddress /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDnsLookup10nameserverEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHostAddress)
	return rv2
}

// /usr/include/qt/QtNetwork/qdnslookup.h:239
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNameserver(const QHostAddress &)

/*

 */
func (this *QDnsLookup) SetNameserver(nameserver QHostAddress_ITF) {
	var convArg0 unsafe.Pointer
	if nameserver != nil && nameserver.QHostAddress_PTR() != nil {
		convArg0 = nameserver.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDnsLookup13setNameserverERK12QHostAddress", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:251
// index:0
// Public Visibility=Default Availability=Available
// [-2] void abort()

/*
Aborts the DNS lookup operation.

If the lookup is already finished, does nothing.
*/
func (this *QDnsLookup) Abort() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDnsLookup5abortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:252
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lookup()

/*
Performs the DNS lookup.

The finished() signal is emitted upon completion.
*/
func (this *QDnsLookup) Lookup() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDnsLookup6lookupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:255
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finished()

/*
This signal is emitted when the reply has finished processing.

Note: Notifier signal for property error. Notifier signal for property errorString.
*/
func (this *QDnsLookup) Finished() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDnsLookup8finishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:256
// index:0
// Public Visibility=Default Availability=Available
// [-2] void nameChanged(const QString &)

/*
This signal is emitted when the lookup name changes. name is the new lookup name.

Note: Notifier signal for property name.
*/
func (this *QDnsLookup) NameChanged(name string) {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDnsLookup11nameChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:257
// index:0
// Public Visibility=Default Availability=Available
// [-2] void typeChanged(QDnsLookup::Type)

/*
This signal is emitted when the lookup type changes. type is the new lookup type.

Note: Notifier signal for property type.
*/
func (this *QDnsLookup) TypeChanged(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDnsLookup11typeChangedENS_4TypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:258
// index:0
// Public Visibility=Default Availability=Available
// [-2] void nameserverChanged(const QHostAddress &)

/*

 */
func (this *QDnsLookup) NameserverChanged(nameserver QHostAddress_ITF) {
	var convArg0 unsafe.Pointer
	if nameserver != nil && nameserver.QHostAddress_PTR() != nil {
		convArg0 = nameserver.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDnsLookup17nameserverChangedERK12QHostAddress", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
Indicates all possible error conditions found during the processing of the DNS lookup.


*/
type QDnsLookup__Error = int

// no error condition.
const QDnsLookup__NoError QDnsLookup__Error = 0

// there was an error initializing the system's DNS resolver.
const QDnsLookup__ResolverError QDnsLookup__Error = 1

// the lookup was aborted using the abort() method.
const QDnsLookup__OperationCancelledError QDnsLookup__Error = 2

// the requested DNS lookup was invalid.
const QDnsLookup__InvalidRequestError QDnsLookup__Error = 3

// the reply returned by the server was invalid.
const QDnsLookup__InvalidReplyError QDnsLookup__Error = 4

// the server encountered an internal failure while processing the request (SERVFAIL).
const QDnsLookup__ServerFailureError QDnsLookup__Error = 5

// the server refused to process the request for security or policy reasons (REFUSED).
const QDnsLookup__ServerRefusedError QDnsLookup__Error = 6

// the requested domain name does not exist (NXDOMAIN).
const QDnsLookup__NotFoundError QDnsLookup__Error = 7

func (this *QDnsLookup) ErrorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QDnsLookup_ErrorItemName(val int) string {
	var nilthis *QDnsLookup
	return nilthis.ErrorItemName(val)
}

/*
Indicates the type of DNS lookup that was performed.


*/
type QDnsLookup__Type = int

//
const QDnsLookup__A QDnsLookup__Type = 1

//
const QDnsLookup__AAAA QDnsLookup__Type = 28

//
const QDnsLookup__ANY QDnsLookup__Type = 255

// canonical name records.
const QDnsLookup__CNAME QDnsLookup__Type = 5

//
const QDnsLookup__MX QDnsLookup__Type = 15

// name server records.
const QDnsLookup__NS QDnsLookup__Type = 2

//
const QDnsLookup__PTR QDnsLookup__Type = 12

//
const QDnsLookup__SRV QDnsLookup__Type = 33

//
const QDnsLookup__TXT QDnsLookup__Type = 16

func (this *QDnsLookup) TypeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QDnsLookup_TypeItemName(val int) string {
	var nilthis *QDnsLookup
	return nilthis.TypeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11399() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
