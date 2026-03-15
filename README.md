# go-musthave-shortener-tpl

Шаблон репозитория для трека «Сервис сокращения URL».

## Начало работы

1. Склонируйте репозиторий в любую подходящую директорию на вашем компьютере.
2. В корне репозитория выполните команду `go mod init <name>` (где `<name>` — адрес вашего репозитория на GitHub без префикса `https://`) для создания модуля.

## Обновление шаблона

Чтобы иметь возможность получать обновления автотестов и других частей шаблона, выполните команду:

```
git remote add -m v2 template https://github.com/Yandex-Practicum/go-musthave-shortener-tpl.git
```

Для обновления кода автотестов выполните команду:

```
git fetch template && git checkout template/v2 .github
```

Затем добавьте полученные изменения в свой репозиторий.

## Запуск автотестов

Для успешного запуска автотестов называйте ветки `iter<number>`, где `<number>` — порядковый номер инкремента. Например, в ветке с названием `iter4` запустятся автотесты для инкрементов с первого по четвёртый.

При мёрже ветки с инкрементом в основную ветку `main` будут запускаться все автотесты.

Подробнее про локальный и автоматический запуск читайте в [README автотестов](https://github.com/Yandex-Practicum/go-autotests).

## Структура проекта

Приведённая в этом репозитории структура проекта является рекомендуемой, но не обязательной.

Это лишь пример организации кода, который поможет вам в реализации сервиса.

При необходимости можно вносить изменения в структуру проекта, использовать любые библиотеки и предпочитаемые структурные паттерны организации кода приложения, например:
- **DDD** (Domain-Driven Design)
- **Clean Architecture**
- **Hexagonal Architecture**
- **Layered Architecture**

File: handler.test.exe
Build ID: O:\GO\shortener\shortener\handler.test.exe2026-03-15 20:06:58.9150417 +0300 MSK
Type: cpu
Time: 2026-03-15 20:06:35 MSK
Duration: 47.80s, Total samples = 19.89s (41.61%)
Showing nodes accounting for -0.61s, 3.07% of 19.89s total
Dropped 14 nodes (cum <= 0.10s)
      flat  flat%   sum%        cum   cum%
     0.58s  2.92%  2.92%      0.58s  2.92%  runtime.stdcall1
    -0.29s  1.46%  1.46%     -0.29s  1.46%  runtime.stdcall2
    -0.28s  1.41%  0.05%     -0.26s  1.31%  runtime.stdcall6
    -0.21s  1.06%  1.01%     -0.21s  1.06%  internal/runtime/atomic.(*Uint32).Add (inline)
    -0.20s  1.01%  2.01%     -0.14s   0.7%  runtime.pcvalue
    -0.17s  0.85%  2.87%     -0.20s  1.01%  runtime.findfunc
    -0.11s  0.55%  3.42%     -0.11s  0.55%  internal/runtime/atomic.(*Uint64).CompareAndSwap (inline)
    -0.10s   0.5%  3.92%     -0.10s   0.5%  runtime.(*mspan).base (inline)
    -0.10s   0.5%  4.42%     -0.10s   0.5%  runtime.(*mspan).heapBitsSmallForAddr
     0.10s   0.5%  3.92%      0.10s   0.5%  runtime.memclrNoHeapPointers
    -0.10s   0.5%  4.42%     -0.02s   0.1%  runtime.unlock2
     0.09s  0.45%  3.97%     -0.06s   0.3%  runtime.(*mcentral).cacheSpan
     0.09s  0.45%  3.52%      0.13s  0.65%  runtime.spanOfUnchecked (inline)
    -0.08s   0.4%  3.92%     -0.12s   0.6%  runtime.(*mcache).releaseAll
    -0.07s  0.35%  4.27%     -0.07s  0.35%  runtime.(*mspan).writeHeapBitsSmall
     0.07s  0.35%  3.92%      0.10s   0.5%  runtime.findObject
     0.07s  0.35%  3.57%      0.04s   0.2%  runtime.gcDrain
     0.06s   0.3%  3.27%      0.09s  0.45%  runtime.(*gcControllerState).trigger
    -0.06s   0.3%  3.57%     -0.06s   0.3%  runtime.mget (inline)
     0.06s   0.3%  3.27%      0.06s   0.3%  runtime.pMask.read (inline)
     0.06s   0.3%  2.97%      0.07s  0.35%  runtime.pageIndexOf (inline)
     0.05s  0.25%  2.71%      0.07s  0.35%  github.com/golang/mock/gomock.(*Call).matches
    -0.05s  0.25%  2.97%     -0.05s  0.25%  internal/runtime/atomic.(*Uint64).Add (inline)
     0.05s  0.25%  2.71%      0.05s  0.25%  internal/runtime/atomic.(*UnsafePointer).Load (inline)
     0.05s  0.25%  2.46%      0.05s  0.25%  runtime.(*heapStatsDelta).merge (inline)
     0.05s  0.25%  2.21%      0.05s  0.25%  runtime.(*timeHistogram).record
     0.05s  0.25%  1.96%     -0.06s   0.3%  runtime.callers.func1
    -0.05s  0.25%  2.21%     -0.05s  0.25%  runtime.nanotime1
    -0.05s  0.25%  2.46%     -0.17s  0.85%  runtime.readmemstats_m
    -0.05s  0.25%  2.71%     -0.15s  0.75%  runtime.stopTheWorldWithSema
     0.05s  0.25%  2.46%     -0.04s   0.2%  runtime.systemstack
     0.04s   0.2%  2.26%      0.04s   0.2%  fmt.(*fmt).fmtInteger
    -0.04s   0.2%  2.46%     -0.04s   0.2%  internal/runtime/atomic.(*Uint32).Load (inline)
    -0.04s   0.2%  2.66%     -0.04s   0.2%  net/url.escape
     0.04s   0.2%  2.46%     -0.02s   0.1%  runtime.(*consistentHeapStats).acquire
     0.04s   0.2%  2.26%      0.04s   0.2%  runtime.(*gcBits).bitp (inline)
    -0.04s   0.2%  2.46%     -0.04s   0.2%  runtime.(*mspan).nextFreeIndex
     0.04s   0.2%  2.26%      0.04s   0.2%  runtime.arenaIdx.l1 (inline)
    -0.04s   0.2%  2.46%     -0.04s   0.2%  runtime.funcInfo.valid (inline)
    -0.04s   0.2%  2.66%      0.05s  0.25%  runtime.gcTrigger.test
    -0.04s   0.2%  2.87%     -0.23s  1.16%  runtime.mallocgc
    -0.04s   0.2%  3.07%     -0.04s   0.2%  runtime.nilinterhash
     0.04s   0.2%  2.87%      0.04s   0.2%  runtime.procyield
     0.04s   0.2%  2.66%      0.04s   0.2%  runtime.readvarint (inline)
    -0.04s   0.2%  2.87%     -0.07s  0.35%  runtime.scanblock
     0.04s   0.2%  2.66%      0.04s   0.2%  runtime.spanOf (inline)
     0.04s   0.2%  2.46%      0.04s   0.2%  runtime.stdcall3
     0.04s   0.2%  2.26%      0.05s  0.25%  sync.(*Pool).Get
     0.04s   0.2%  2.06%     -0.01s  0.05%  type:.hash.github.com/golang/mock/gomock.callSetKey
    -0.03s  0.15%  2.21%     -0.07s  0.35%  github.com/FedorSidorow/shortener/internal/handler.(*APIHandler).GenerateShortKeyHandler
     0.03s  0.15%  2.06%      0.03s  0.15%  internal/bytealg.IndexByteString
    -0.03s  0.15%  2.21%     -0.03s  0.15%  internal/sync.(*Mutex).Unlock (inline)
    -0.03s  0.15%  2.36%     -0.03s  0.15%  runtime.(*gcWork).putObjFast (inline)
     0.03s  0.15%  2.21%      0.06s   0.3%  runtime.(*inlineUnwinder).next
     0.03s  0.15%  2.06%     -0.04s   0.2%  runtime.(*limiterEvent).stop
    -0.03s  0.15%  2.21%     -0.03s  0.15%  runtime.(*mspan).refillAllocCache
     0.03s  0.15%  2.06%     -0.07s  0.35%  runtime.(*mspan).typePointersOfUnchecked
    -0.03s  0.15%  2.21%     -0.12s   0.6%  runtime.(*spanSet).pop
    -0.03s  0.15%  2.36%     -0.02s   0.1%  runtime.(*spanSet).push
    -0.03s  0.15%  2.51%     -0.03s  0.15%  runtime.findmoduledatap (inline)
    -0.03s  0.15%  2.66%     -0.05s  0.25%  runtime.funcInfo.srcFunc (inline)
    -0.03s  0.15%  2.82%     -0.03s  0.15%  runtime.funcdata (inline)
    -0.03s  0.15%  2.97%      0.44s  2.21%  runtime.notewakeup
     0.03s  0.15%  2.82%      0.03s  0.15%  runtime.stdcall4
     0.03s  0.15%  2.66%      0.03s  0.15%  sync/atomic.(*Int32).Add (inline)
     0.02s   0.1%  2.56%      0.04s   0.2%  bytes.(*Buffer).grow
     0.02s   0.1%  2.46%      0.07s  0.35%  fmt.newPrinter
     0.02s   0.1%  2.36%      0.02s   0.1%  gogo
    -0.02s   0.1%  2.46%     -0.02s   0.1%  internal/runtime/atomic.(*Uint8).Load (inline)
    -0.02s   0.1%  2.56%     -0.04s   0.2%  internal/runtime/maps.(*Iter).Next
     0.02s   0.1%  2.46%     -0.01s  0.05%  internal/runtime/maps.(*Map).putSlotSmall
     0.02s   0.1%  2.36%      0.02s   0.1%  internal/sync.(*Mutex).Lock (inline)
    -0.02s   0.1%  2.46%     -0.22s  1.11%  net/http.NewRequestWithContext
    -0.02s   0.1%  2.56%     -0.06s   0.3%  net/url.(*URL).setPath
    -0.02s   0.1%  2.66%     -0.03s  0.15%  net/url.parseHost
    -0.02s   0.1%  2.77%     -0.02s   0.1%  reflect.(*rtype).Kind (partial-inline)
    -0.02s   0.1%  2.87%     -0.04s   0.2%  reflect.deepValueEqual
     0.02s   0.1%  2.77%      0.02s   0.1%  reflect.directlyAssignable
     0.02s   0.1%  2.66%      0.02s   0.1%  reflect.packEface
    -0.02s   0.1%  2.77%     -0.11s  0.55%  runtime.(*Frames).Next
     0.02s   0.1%  2.66%      0.02s   0.1%  runtime.(*gcControllerState).heapGoalInternal
    -0.02s   0.1%  2.77%     -0.02s   0.1%  runtime.(*mspan).init
     0.02s   0.1%  2.66%      0.06s   0.3%  runtime.(*mspan).markBitsForIndex (inline)
     0.02s   0.1%  2.56%      0.03s  0.15%  runtime.(*randomOrder).reset (inline)
     0.02s   0.1%  2.46%      0.03s  0.15%  runtime.(*sweepLocked).sweep
    -0.02s   0.1%  2.56%     -0.02s   0.1%  runtime.(*sysMemStat).add
    -0.02s   0.1%  2.66%     -0.14s   0.7%  runtime.(*unwinder).initAt
     0.02s   0.1%  2.56%     -0.03s  0.15%  runtime.callers
    -0.02s   0.1%  2.66%     -0.02s   0.1%  runtime.cansemacquire (inline)
    -0.02s   0.1%  2.77%     -0.02s   0.1%  runtime.cheaprand (inline)
     0.02s   0.1%  2.66%      0.01s  0.05%  runtime.cheaprandn (inline)
    -0.02s   0.1%  2.77%     -0.02s   0.1%  runtime.divRoundUp (inline)
     0.02s   0.1%  2.66%      0.02s   0.1%  runtime.funcInfo.entry (inline)
    -0.02s   0.1%  2.77%     -0.02s   0.1%  runtime.getMCache (inline)
    -0.02s   0.1%  2.87%     -0.09s  0.45%  runtime.heapSetTypeNoHeader (inline)
    -0.02s   0.1%  2.97%     -0.02s   0.1%  runtime.limiterEventStamp.duration (inline)
    -0.02s   0.1%  3.07%     -0.07s  0.35%  runtime.makeslice
     0.02s   0.1%  2.97%      0.07s  0.35%  runtime.mallocgcSmallNoscan
    -0.02s   0.1%  3.07%     -0.15s  0.75%  runtime.mapassign_faststr
    -0.02s   0.1%  3.17%     -0.02s   0.1%  runtime.newMarkBits
    -0.02s   0.1%  3.27%     -0.14s   0.7%  runtime.newobject
     0.02s   0.1%  3.17%      0.02s   0.1%  runtime.nextFreeFast (inline)
     0.02s   0.1%  3.07%      0.02s   0.1%  runtime.pidleput
     0.02s   0.1%  2.97%     -0.02s   0.1%  runtime.preemptall
     0.02s   0.1%  2.87%      0.02s   0.1%  runtime.rand
     0.02s   0.1%  2.77%      0.02s   0.1%  runtime.releasem (inline)
     0.02s   0.1%  2.66%      0.10s   0.5%  runtime.scanobject
    -0.02s   0.1%  2.77%      0.54s  2.71%  runtime.semawakeup
    -0.02s   0.1%  2.87%     -0.05s  0.25%  runtime.semrelease1
     0.02s   0.1%  2.77%     -0.04s   0.2%  runtime.stackcache_clear
    -0.02s   0.1%  2.87%     -0.02s   0.1%  runtime.stdcall0
     0.02s   0.1%  2.77%      0.06s   0.3%  runtime.step
     0.02s   0.1%  2.66%      0.09s  0.45%  runtime.sweepone
     0.02s   0.1%  2.56%      0.03s  0.15%  runtime.tracebackPCs
     0.02s   0.1%  2.46%      0.02s   0.1%  runtime.typePointers.nextFast (inline)
     0.01s  0.05%  2.41%      0.01s  0.05%  aeshashbody
    -0.01s  0.05%  2.46%     -0.01s  0.05%  context.(*valueCtx).Value
    -0.01s  0.05%  2.51%     -0.02s   0.1%  context.WithValue
     0.01s  0.05%  2.46%      0.01s  0.05%  fmt.(*buffer).writeString (inline)
     0.01s  0.05%  2.41%     -0.01s  0.05%  fmt.(*fmt).fmtS
    -0.01s  0.05%  2.46%     -0.02s   0.1%  fmt.(*fmt).padString
     0.01s  0.05%  2.41%      0.05s  0.25%  fmt.(*pp).doPrintf
    -0.01s  0.05%  2.46%     -0.04s   0.2%  github.com/FedorSidorow/shortener/internal/handler.(*MockShortenerServicer).GenerateShortURL       
     0.01s  0.05%  2.41%      0.19s  0.96%  github.com/FedorSidorow/shortener/internal/handler.(*MockShortenerServicer).GetURLByKey
    -0.01s  0.05%  2.46%      0.02s   0.1%  github.com/FedorSidorow/shortener/internal/handler.(*MockShortenerServicerMockRecorder).GenerateShortURL
    -0.01s  0.05%  2.51%     -0.42s  2.11%  github.com/FedorSidorow/shortener/internal/handler.BenchmarkAPIHandler_GetURLByKeyHandler.func1    
     0.01s  0.05%  2.46%      0.01s  0.05%  github.com/go-chi/chi/v5.(*Context).URLParam (inline)
     0.01s  0.05%  2.41%      0.01s  0.05%  github.com/go-chi/chi/v5.RouteContext (inline)
     0.01s  0.05%  2.36%     -0.01s  0.05%  github.com/golang/mock/gomock.(*Controller).Call.func1
     0.01s  0.05%  2.31%      0.13s  0.65%  github.com/golang/mock/gomock.(*Controller).RecordCallWithMethodType
     0.01s  0.05%  2.26%     -0.02s   0.1%  github.com/golang/mock/gomock.callSet.Remove (inline)
     0.01s  0.05%  2.21%      0.07s  0.35%  github.com/golang/mock/gomock.callerInfo
    -0.01s  0.05%  2.26%      0.14s   0.7%  github.com/golang/mock/gomock.newCall
     0.01s  0.05%  2.21%     -0.16s   0.8%  github.com/stretchr/testify/require.NoError
    -0.01s  0.05%  2.26%     -0.02s   0.1%  internal/abi.(*FuncType).OutSlice (inline)
    -0.01s  0.05%  2.31%     -0.01s  0.05%  internal/abi.(*Type).IfaceIndir (inline)
    -0.01s  0.05%  2.36%     -0.01s  0.05%  internal/abi.(*Type).Size (inline)
    -0.01s  0.05%  2.41%     -0.01s  0.05%  internal/abi.addChecked (inline)
     0.01s  0.05%  2.36%      0.01s  0.05%  internal/bytealg.LastIndexByteString (inline)
     0.01s  0.05%  2.31%      0.01s  0.05%  internal/chacha8rand.(*State).Next (inline)
    -0.01s  0.05%  2.36%     -0.01s  0.05%  internal/chacha8rand.(*State).Refill
     0.01s  0.05%  2.31%      0.01s  0.05%  internal/runtime/atomic.(*Int32).CompareAndSwap (inline)
     0.01s  0.05%  2.26%      0.01s  0.05%  internal/runtime/atomic.(*Uint32).CompareAndSwap (inline)
    -0.01s  0.05%  2.31%     -0.01s  0.05%  internal/runtime/atomic.(*Uint8).Store (inline)
    -0.01s  0.05%  2.36%     -0.01s  0.05%  internal/runtime/maps.(*Map).getWithKeySmall
    -0.01s  0.05%  2.41%     -0.01s  0.05%  internal/runtime/maps.(*ctrlGroup).setEmpty (inline)
    -0.01s  0.05%  2.46%     -0.01s  0.05%  internal/runtime/maps.(*groupReference).elem (inline)
    -0.01s  0.05%  2.51%     -0.01s  0.05%  internal/runtime/maps.(*groupReference).key (inline)
    -0.01s  0.05%  2.56%     -0.01s  0.05%  internal/runtime/maps.(*groupsReference).group (inline)
     0.01s  0.05%  2.51%     -0.02s   0.1%  internal/runtime/maps.NewMap
     0.01s  0.05%  2.46%      0.01s  0.05%  internal/runtime/maps.h1 (inline)
    -0.01s  0.05%  2.51%     -0.01s  0.05%  internal/runtime/maps.h2 (inline)
    -0.01s  0.05%  2.56%     -0.02s   0.1%  internal/stringslite.HasSuffix (inline)
    -0.01s  0.05%  2.61%     -0.01s  0.05%  memeqbody
    -0.01s  0.05%  2.66%     -0.01s  0.05%  net/http.(*Request).Context (inline)
    -0.01s  0.05%  2.71%     -0.11s  0.55%  net/http.Header.Clone (inline)
    -0.01s  0.05%  2.77%     -0.01s  0.05%  net/http/httptest.(*ResponseRecorder).Header
    -0.01s  0.05%  2.82%     -0.13s  0.65%  net/http/httptest.(*ResponseRecorder).WriteHeader
    -0.01s  0.05%  2.87%     -0.01s  0.05%  net/http/httptest.checkWriteHeaderCode (inline)
    -0.01s  0.05%  2.92%     -0.02s   0.1%  net/textproto.canonicalMIMEHeaderKey
    -0.01s  0.05%  2.97%     -0.18s   0.9%  net/url.Parse
     0.01s  0.05%  2.92%     -0.17s  0.85%  net/url.parse
     0.01s  0.05%  2.87%      0.04s   0.2%  reflect.(*rtype).AssignableTo
     0.01s  0.05%  2.82%      0.02s   0.1%  reflect.(*rtype).IsVariadic
    -0.01s  0.05%  2.87%     -0.04s   0.2%  reflect.(*rtype).Out
     0.01s  0.05%  2.82%      0.01s  0.05%  reflect.(*rtype).common
    -0.01s  0.05%  2.87%     -0.01s  0.05%  reflect.TypeOf (inline)
    -0.01s  0.05%  2.92%     -0.01s  0.05%  reflect.Value.lenNonSlice
    -0.01s  0.05%  2.97%     -0.03s  0.15%  reflect.ValueOf (inline)
     0.01s  0.05%  2.92%      0.01s  0.05%  reflect.add (inline)
    -0.01s  0.05%  2.97%     -0.01s  0.05%  reflect.flag.kind (inline)
     0.01s  0.05%  2.92%      0.01s  0.05%  reflect.packIfaceValueIntoEmptyIface (inline)
    -0.01s  0.05%  2.97%     -0.01s  0.05%  reflect.toType (inline)
    -0.01s  0.05%  3.02%     -0.02s   0.1%  reflect.unpackEface (inline)
    -0.01s  0.05%  3.07%      0.01s  0.05%  reflect.valueInterface
     0.01s  0.05%  3.02%      0.03s  0.15%  runtime.(*activeSweep).begin (inline)
     0.01s  0.05%  2.97%     -0.02s   0.1%  runtime.(*activeSweep).end
    -0.01s  0.05%  3.02%     -0.06s   0.3%  runtime.(*consistentHeapStats).release
     0.01s  0.05%  2.97%      0.01s  0.05%  runtime.(*fixalloc).alloc
    -0.01s  0.05%  3.02%     -0.01s  0.05%  runtime.(*gcCPULimiterState).accumulate
     0.01s  0.05%  2.97%     -0.20s  1.01%  runtime.(*gcControllerState).enlistWorker
    -0.01s  0.05%  3.02%     -0.01s  0.05%  runtime.(*gcWork).tryGetObjFast (inline)
     0.01s  0.05%  2.97%      0.01s  0.05%  runtime.(*guintptr).set (inline)
    -0.01s  0.05%  3.02%      0.10s   0.5%  runtime.(*inlineUnwinder).resolveInternal (inline)
     0.01s  0.05%  2.97%     -0.04s   0.2%  runtime.(*inlineUnwinder).srcFunc (inline)
    -0.01s  0.05%  3.02%     -0.01s  0.05%  runtime.(*lfstack).pop (inline)
     0.01s  0.05%  2.97%      0.01s  0.05%  runtime.(*m).clearAllpSnapshot
    -0.01s  0.05%  3.02%     -0.17s  0.85%  runtime.(*mcache).nextFree
    -0.01s  0.05%  3.07%     -0.12s   0.6%  runtime.(*mcache).refill
     0.01s  0.05%  3.02%      0.01s  0.05%  runtime.(*mheap).allocNeedsZero
     0.01s  0.05%  2.97%      0.03s  0.15%  runtime.(*mheap).allocSpan
    -0.01s  0.05%  3.02%      0.01s  0.05%  runtime.(*mheap).freeSpan (inline)
     0.01s  0.05%  2.97%      0.01s  0.05%  runtime.(*mheap).freeSpanLocked
    -0.01s  0.05%  3.02%     -0.02s   0.1%  runtime.(*mheap).initSpan
    -0.01s  0.05%  3.07%     -0.01s  0.05%  runtime.(*moduledata).textOff (inline)
     0.01s  0.05%  3.02%      0.01s  0.05%  runtime.(*mspan).countAlloc (inline)
     0.01s  0.05%  2.97%      0.03s  0.15%  runtime.(*pageAlloc).update
     0.01s  0.05%  2.92%      0.01s  0.05%  runtime.(*pallocBits).summarize
    -0.01s  0.05%  2.97%     -0.01s  0.05%  runtime.(*randomEnum).next (inline)
     0.01s  0.05%  2.92%      0.01s  0.05%  runtime.(*sweepLocker).tryAcquire
     0.01s  0.05%  2.87%      0.02s   0.1%  runtime.(*timers).run
     0.01s  0.05%  2.82%     -0.05s  0.25%  runtime.(*unwinder).next
    -0.01s  0.05%  2.87%     -0.01s  0.05%  runtime.(*unwinder).symPC
    -0.01s  0.05%  2.92%     -0.39s  1.96%  runtime.ReadMemStats
     0.01s  0.05%  2.87%     -0.16s   0.8%  runtime.ReadMemStats.func1
     0.01s  0.05%  2.82%      0.01s  0.05%  runtime.add (inline)
    -0.01s  0.05%  2.87%     -0.01s  0.05%  runtime.adjustframe
     0.01s  0.05%  2.82%      0.01s  0.05%  runtime.arenaIndex (inline)
    -0.01s  0.05%  2.87%     -0.01s  0.05%  runtime.cgocall
    -0.01s  0.05%  2.92%     -0.01s  0.05%  runtime.checkRunqsNoP
     0.01s  0.05%  2.87%      0.02s   0.1%  runtime.deductSweepCredit
    -0.01s  0.05%  2.92%     -0.01s  0.05%  runtime.duffcopy
     0.01s  0.05%  2.87%      0.01s  0.05%  runtime.duffzero
     0.01s  0.05%  2.82%      0.02s   0.1%  runtime.findnull
     0.01s  0.05%  2.77%     -0.16s   0.8%  runtime.flushallmcaches
    -0.01s  0.05%  2.82%     -0.17s  0.85%  runtime.flushmcache
    -0.01s  0.05%  2.87%      0.01s  0.05%  runtime.funcfile
     0.01s  0.05%  2.82%     -0.09s  0.45%  runtime.funcline1
     0.01s  0.05%  2.77%     -0.12s   0.6%  runtime.funcspdelta (inline)
    -0.01s  0.05%  2.82%     -0.01s  0.05%  runtime.gFromSP (inline)
     0.01s  0.05%  2.77%      0.01s  0.05%  runtime.gcBgMarkWorker.func1
     0.01s  0.05%  2.71%      0.01s  0.05%  runtime.gcd (inline)
    -0.01s  0.05%  2.77%     -0.01s  0.05%  runtime.globrunqgetbatch
    -0.01s  0.05%  2.82%      0.09s  0.45%  runtime.goschedImpl
    -0.01s  0.05%  2.87%     -0.12s   0.6%  runtime.greyobject
     0.01s  0.05%  2.82%      0.01s  0.05%  runtime.headTailIndex.head (inline)
     0.01s  0.05%  2.77%      0.01s  0.05%  runtime.ifaceeq
    -0.01s  0.05%  2.82%     -0.01s  0.05%  runtime.interhash
     0.01s  0.05%  2.77%      0.01s  0.05%  runtime.isDirectIface (inline)
    -0.01s  0.05%  2.82%     -0.01s  0.05%  runtime.lfstackPack (inline)
     0.01s  0.05%  2.77%      0.01s  0.05%  runtime.limiterEventStamp.typ (inline)
     0.01s  0.05%  2.71%     -0.11s  0.55%  runtime.lock2
     0.01s  0.05%  2.66%      0.01s  0.05%  runtime.mProf_Free
    -0.01s  0.05%  2.71%     -0.01s  0.05%  runtime.makeSpanClass (inline)
    -0.01s  0.05%  2.77%     -0.22s  1.11%  runtime.mallocgcSmallScanNoHeader
     0.01s  0.05%  2.71%     -0.01s  0.05%  runtime.mallocgcTiny
    -0.01s  0.05%  2.77%     -0.05s  0.25%  runtime.mapIterStart
    -0.01s  0.05%  2.82%     -0.01s  0.05%  runtime.mapaccess1_faststr
    -0.01s  0.05%  2.87%     -0.01s  0.05%  runtime.markBits.isMarked (inline)
    -0.01s  0.05%  2.92%     -0.02s   0.1%  runtime.mcall
    -0.01s  0.05%  2.97%     -0.01s  0.05%  runtime.memequal
    -0.01s  0.05%  3.02%     -0.01s  0.05%  runtime.memhash
     0.01s  0.05%  2.97%      0.01s  0.05%  runtime.memmove
     0.01s  0.05%  2.92%      0.01s  0.05%  runtime.mergeSummaries
    -0.01s  0.05%  2.97%      0.06s   0.3%  runtime.morestack
     0.01s  0.05%  2.92%      0.01s  0.05%  runtime.mutexSampleContention (inline)
    -0.01s  0.05%  2.97%     -0.12s   0.6%  runtime.newarray
     0.01s  0.05%  2.92%      0.01s  0.05%  runtime.notetsleep_internal
     0.01s  0.05%  2.87%      0.01s  0.05%  runtime.pMask.clear (inline)
    -0.01s  0.05%  2.92%     -0.01s  0.05%  runtime.pMask.set (inline)
     0.01s  0.05%  2.87%      0.01s  0.05%  runtime.pcdatastart (inline)
    -0.01s  0.05%  2.92%     -0.05s  0.25%  runtime.pidleget
     0.01s  0.05%  2.87%     -0.25s  1.26%  runtime.preemptone
    -0.01s  0.05%  2.92%      0.03s  0.15%  runtime.procresize
    -0.01s  0.05%  2.97%     -0.01s  0.05%  runtime.releasep
    -0.01s  0.05%  3.02%      0.03s  0.15%  runtime.resetspinning
    -0.01s  0.05%  3.07%     -0.01s  0.05%  runtime.roundupsize (inline)
    -0.01s  0.05%  3.12%     -0.01s  0.05%  runtime.runqempty (inline)
    -0.01s  0.05%  3.17%     -0.01s  0.05%  runtime.semacreate
     0.01s  0.05%  3.12%      0.01s  0.05%  runtime.spanClass.sizeclass (inline)
    -0.01s  0.05%  3.17%     -0.01s  0.05%  runtime.spanHasSpecials (inline)
     0.01s  0.05%  3.12%      0.02s   0.1%  runtime.stdcall
    -0.01s  0.05%  3.17%     -0.01s  0.05%  runtime.stdcall7
    -0.01s  0.05%  3.22%     -0.01s  0.05%  runtime.stkbucket
    -0.01s  0.05%  3.27%     -0.17s  0.85%  runtime.stopTheWorld
     0.01s  0.05%  3.22%      0.03s  0.15%  runtime.stopm
    -0.01s  0.05%  3.27%     -0.01s  0.05%  runtime.strhash
     0.01s  0.05%  3.22%      0.01s  0.05%  runtime.traceAcquire (inline)
    -0.01s  0.05%  3.27%     -0.01s  0.05%  runtime.traceLocker.ok (inline)
     0.01s  0.05%  3.22%      0.01s  0.05%  runtime.traceRelease (inline)
     0.01s  0.05%  3.17%      0.01s  0.05%  runtime.waitReason.isWaitingForSuspendG (inline)
    -0.01s  0.05%  3.22%      0.05s  0.25%  sync.(*RWMutex).Lock
     0.01s  0.05%  3.17%      0.01s  0.05%  sync.(*RWMutex).Unlock
     0.01s  0.05%  3.12%      0.01s  0.05%  testing.(*B).Helper
     0.01s  0.05%  3.07%      0.05s  0.25%  testing.(*common).Helper
    -0.01s  0.05%  3.12%     -0.01s  0.05%  testing.highPrecisionTime.sub
     0.01s  0.05%  3.07%     -0.01s  0.05%  type:.hash.reflect.visit
         0     0%  3.07%      0.04s   0.2%  bytes.(*Buffer).Write
         0     0%  3.07%      0.04s   0.2%  fmt.(*pp).fmtInteger
         0     0%  3.07%     -0.01s  0.05%  fmt.(*pp).fmtString
         0     0%  3.07%     -0.01s  0.05%  fmt.(*pp).free
         0     0%  3.07%      0.03s  0.15%  fmt.(*pp).printArg
         0     0%  3.07%      0.16s   0.8%  fmt.Sprintf
         0     0%  3.07%      0.01s  0.05%  gcWriteBarrier
         0     0%  3.07%     -0.01s  0.05%  github.com/FedorSidorow/shortener/internal/auth.UserIDFrom (inline)
         0     0%  3.07%     -0.01s  0.05%  github.com/FedorSidorow/shortener/internal/auth.WithUserID (inline)
         0     0%  3.07%      0.03s  0.15%  github.com/FedorSidorow/shortener/internal/handler.(*APIHandler).GetURLByKeyHandler
         0     0%  3.07%      0.18s   0.9%  github.com/FedorSidorow/shortener/internal/handler.(*MockShortenerServicerMockRecorder).GetURLByKey
         0     0%  3.07%     -0.46s  2.31%  github.com/FedorSidorow/shortener/internal/handler.BenchmarkAPIHandler_GenerateShortkeyHandler.func1
         0     0%  3.07%     -0.01s  0.05%  github.com/FedorSidorow/shortener/internal/handler.TestAPIHandler_GenerateShortKeyHandler
         0     0%  3.07%     -0.01s  0.05%  github.com/FedorSidorow/shortener/internal/storage/inmemorystore.NewStorage
         0     0%  3.07%     -0.03s  0.15%  github.com/go-chi/chi/v5.(*RouteParams).Add (inline)
         0     0%  3.07%     -0.03s  0.15%  github.com/go-chi/chi/v5.NewRouteContext (inline)
         0     0%  3.07%      0.02s   0.1%  github.com/go-chi/chi/v5.URLParam
         0     0%  3.07%      0.03s  0.15%  github.com/golang/mock/gomock.(*Call).addAction (inline)
         0     0%  3.07%      0.01s  0.05%  github.com/golang/mock/gomock.(*Controller).Call
         0     0%  3.07%      0.03s  0.15%  github.com/golang/mock/gomock.callSet.Add (inline)
         0     0%  3.07%      0.09s  0.45%  github.com/golang/mock/gomock.callSet.FindMatch
         0     0%  3.07%     -0.01s  0.05%  github.com/golang/mock/gomock.newCall.func1
         0     0%  3.07%     -0.01s  0.05%  internal/runtime/atomic.(*Bool).Load (inline)
         0     0%  3.07%     -0.01s  0.05%  internal/runtime/atomic.(*Bool).Store (inline)
         0     0%  3.07%     -0.13s  0.65%  internal/runtime/maps.(*Map).growToSmall
         0     0%  3.07%     -0.01s  0.05%  internal/runtime/maps.(*Map).putSlotSmallFastStr
         0     0%  3.07%     -0.05s  0.25%  internal/runtime/maps.NewEmptyMap (inline)
         0     0%  3.07%     -0.12s   0.6%  internal/runtime/maps.newGroups (inline)
         0     0%  3.07%     -0.12s   0.6%  internal/runtime/maps.newarray
         0     0%  3.07%      0.01s  0.05%  internal/runtime/maps.rand
         0     0%  3.07%      0.01s  0.05%  internal/stringslite.IndexByte (inline)
         0     0%  3.07%      0.01s  0.05%  internal/syscall/windows.QueryPerformanceCounter
         0     0%  3.07%     -0.01s  0.05%  internal/syscall/windows/registry.Key.GetMUIStringValue
         0     0%  3.07%     -0.01s  0.05%  internal/syscall/windows/registry.regLoadMUIString
         0     0%  3.07%     -0.06s   0.3%  io.NopCloser (inline)
         0     0%  3.07%     -0.01s  0.05%  log.(*Logger).output
         0     0%  3.07%     -0.01s  0.05%  log.Printf (inline)
         0     0%  3.07%     -0.01s  0.05%  log.formatHeader
         0     0%  3.07%     -0.02s   0.1%  net/http.(*Request).WithContext (inline)
         0     0%  3.07%     -0.12s   0.6%  net/http.Header.Set (inline)
         0     0%  3.07%     -0.22s  1.11%  net/http.NewRequest (inline)
         0     0%  3.07%      0.04s   0.2%  net/http/httptest.(*ResponseRecorder).Write
         0     0%  3.07%     -0.07s  0.35%  net/http/httptest.NewRecorder (inline)
         0     0%  3.07%     -0.12s   0.6%  net/textproto.MIMEHeader.Set (inline)
         0     0%  3.07%     -0.02s   0.1%  net/url.parseAuthority
         0     0%  3.07%     -0.02s   0.1%  reflect.(*rtype).NumOut
         0     0%  3.07%     -0.03s  0.15%  reflect.DeepEqual
         0     0%  3.07%      0.01s  0.05%  reflect.Value.Index
         0     0%  3.07%      0.01s  0.05%  reflect.Value.Interface (inline)
         0     0%  3.07%     -0.01s  0.05%  reflect.Value.Len (inline)
         0     0%  3.07%     -0.01s  0.05%  reflect.Value.pointer (inline)
         0     0%  3.07%     -0.01s  0.05%  reflect.deepValueEqual.func2 (inline)
         0     0%  3.07%     -0.05s  0.25%  runtime.(*atomicHeadTailIndex).cas (inline)
         0     0%  3.07%      0.05s  0.25%  runtime.(*atomicMSpanPointer).Load (inline)
         0     0%  3.07%      0.05s  0.25%  runtime.(*consistentHeapStats).unsafeRead
         0     0%  3.07%     -0.01s  0.05%  runtime.(*gcCPULimiterState).update
         0     0%  3.07%     -0.01s  0.05%  runtime.(*gcCPULimiterState).updateLocked
         0     0%  3.07%     -0.01s  0.05%  runtime.(*gcControllerState).findRunnableGCWorker
         0     0%  3.07%     -0.01s  0.05%  runtime.(*gcControllerState).heapGoal (inline)
         0     0%  3.07%     -0.09s  0.45%  runtime.(*gcWork).balance
         0     0%  3.07%     -0.01s  0.05%  runtime.(*gcWork).init
         0     0%  3.07%     -0.12s   0.6%  runtime.(*gcWork).putObj
         0     0%  3.07%     -0.01s  0.05%  runtime.(*lfstack).push
         0     0%  3.07%      0.01s  0.05%  runtime.(*mcache).prepareForSweep
         0     0%  3.07%      0.03s  0.15%  runtime.(*mcentral).grow
         0     0%  3.07%     -0.02s   0.1%  runtime.(*mcentral).uncacheSpan
         0     0%  3.07%      0.02s   0.1%  runtime.(*mheap).alloc.func1
         0     0%  3.07%      0.01s  0.05%  runtime.(*mheap).allocMSpanLocked
         0     0%  3.07%      0.01s  0.05%  runtime.(*mheap).allocManual
         0     0%  3.07%      0.01s  0.05%  runtime.(*mheap).nextSpanForSweep
         0     0%  3.07%      0.03s  0.15%  runtime.(*mspan).initHeapBits
         0     0%  3.07%      0.03s  0.15%  runtime.(*pageAlloc).free
         0     0%  3.07%      0.04s   0.2%  runtime.(*pageAlloc).scavenge.func1
         0     0%  3.07%      0.04s   0.2%  runtime.(*pageAlloc).scavengeOne
         0     0%  3.07%      0.01s  0.05%  runtime.(*scavengerState).init.func1
         0     0%  3.07%      0.01s  0.05%  runtime.(*scavengerState).sleep
         0     0%  3.07%      0.01s  0.05%  runtime.(*scavengerState).wake
         0     0%  3.07%      0.02s   0.1%  runtime.(*sweepLocked).sweep.(*mheap).freeSpan.func2
         0     0%  3.07%      0.02s   0.1%  runtime.(*timer).maybeAdd
         0     0%  3.07%      0.02s   0.1%  runtime.(*timer).modify
         0     0%  3.07%      0.02s   0.1%  runtime.(*timer).reset (inline)
         0     0%  3.07%      0.01s  0.05%  runtime.(*timer).unlockAndRun
         0     0%  3.07%      0.02s   0.1%  runtime.(*timers).check
         0     0%  3.07%     -0.11s  0.55%  runtime.Caller
         0     0%  3.07%     -0.04s   0.2%  runtime.Callers (inline)
         0     0%  3.07%      0.03s  0.15%  runtime.CallersFrames (inline)
         0     0%  3.07%      0.01s  0.05%  runtime.acquirep
         0     0%  3.07%     -0.01s  0.05%  runtime.addspecial
         0     0%  3.07%      0.01s  0.05%  runtime.allocm
         0     0%  3.07%      0.01s  0.05%  runtime.bgscavenge
         0     0%  3.07%      0.10s   0.5%  runtime.bgsweep
         0     0%  3.07%      0.01s  0.05%  runtime.casGToWaiting (inline)
         0     0%  3.07%      0.02s   0.1%  runtime.casGToWaitingForSuspendG
         0     0%  3.07%      0.01s  0.05%  runtime.checkTimersNoP
         0     0%  3.07%     -0.06s   0.3%  runtime.convT
         0     0%  3.07%      0.01s  0.05%  runtime.convTnoptr
         0     0%  3.07%     -0.01s  0.05%  runtime.copystack
         0     0%  3.07%     -0.03s  0.15%  runtime.deductAssistCredit
         0     0%  3.07%     -0.14s   0.7%  runtime.findRunnable
         0     0%  3.07%     -0.01s  0.05%  runtime.forEachPInternal
         0     0%  3.07%      0.01s  0.05%  runtime.freeSpecial
         0     0%  3.07%      0.01s  0.05%  runtime.funcNameForPrint (inline)
         0     0%  3.07%      0.01s  0.05%  runtime.funcNamePiecesForPrint
         0     0%  3.07%     -0.03s  0.15%  runtime.gcAssistAlloc
         0     0%  3.07%     -0.05s  0.25%  runtime.gcAssistAlloc.func2
         0     0%  3.07%     -0.05s  0.25%  runtime.gcAssistAlloc1
         0     0%  3.07%      0.22s  1.11%  runtime.gcBgMarkWorker
         0     0%  3.07%      0.03s  0.15%  runtime.gcBgMarkWorker.func2
         0     0%  3.07%     -0.07s  0.35%  runtime.gcDrainMarkWorkerDedicated (inline)
         0     0%  3.07%      0.11s  0.55%  runtime.gcDrainMarkWorkerIdle (inline)
         0     0%  3.07%     -0.05s  0.25%  runtime.gcDrainN
         0     0%  3.07%     -0.01s  0.05%  runtime.gcMarkDone.forEachP.func5
         0     0%  3.07%     -0.22s  1.11%  runtime.gcstopm
         0     0%  3.07%      0.01s  0.05%  runtime.getempty.func1
         0     0%  3.07%      0.06s   0.3%  runtime.gopreempt_m (inline)
         0     0%  3.07%      0.03s  0.15%  runtime.gosched_m
         0     0%  3.07%      0.02s   0.1%  runtime.gostringnocopy (inline)
         0     0%  3.07%      0.01s  0.05%  runtime.headTailIndex.split (inline)
         0     0%  3.07%      0.01s  0.05%  runtime.injectglist
         0     0%  3.07%      0.01s  0.05%  runtime.injectglist.func1
         0     0%  3.07%     -0.11s  0.55%  runtime.lock (inline)
         0     0%  3.07%     -0.02s   0.1%  runtime.lock2.osyield.func1
         0     0%  3.07%     -0.11s  0.55%  runtime.lockWithRank (inline)
         0     0%  3.07%     -0.01s  0.05%  runtime.mPark (inline)
         0     0%  3.07%     -0.02s   0.1%  runtime.mProf_Malloc
         0     0%  3.07%     -0.01s  0.05%  runtime.mProf_Malloc.func1
         0     0%  3.07%     -0.02s   0.1%  runtime.makemap
         0     0%  3.07%     -0.05s  0.25%  runtime.makemap_small
         0     0%  3.07%      0.02s   0.1%  runtime.mapaccess2_fast64
         0     0%  3.07%     -0.08s   0.4%  runtime.markroot
         0     0%  3.07%     -0.07s  0.35%  runtime.markrootBlock
         0     0%  3.07%     -0.01s  0.05%  runtime.markrootSpans
         0     0%  3.07%      0.01s  0.05%  runtime.mcommoninit
         0     0%  3.07%     -0.05s  0.25%  runtime.nanotime (inline)
         0     0%  3.07%     -0.27s  1.36%  runtime.netpoll
         0     0%  3.07%      0.01s  0.05%  runtime.netpollBreak
         0     0%  3.07%      0.04s   0.2%  runtime.newInlineUnwinder
         0     0%  3.07%      0.02s   0.1%  runtime.newm
         0     0%  3.07%      0.01s  0.05%  runtime.newm1
         0     0%  3.07%      0.01s  0.05%  runtime.newosproc
         0     0%  3.07%      0.05s  0.25%  runtime.newstack
         0     0%  3.07%     -0.01s  0.05%  runtime.notesleep
         0     0%  3.07%      0.01s  0.05%  runtime.notetsleep
         0     0%  3.07%     -0.02s   0.1%  runtime.osyield (inline)
         0     0%  3.07%     -0.04s   0.2%  runtime.park_m
         0     0%  3.07%      0.11s  0.55%  runtime.pcdatavalue1
         0     0%  3.07%      0.02s   0.1%  runtime.pidlegetSpinning
         0     0%  3.07%     -0.26s  1.31%  runtime.preemptM
         0     0%  3.07%     -0.02s   0.1%  runtime.profilealloc
         0     0%  3.07%     -0.01s  0.05%  runtime.putfull
         0     0%  3.07%      0.03s  0.15%  runtime.rawbyteslice
         0     0%  3.07%      0.01s  0.05%  runtime.resetForSleep
         0     0%  3.07%     -0.02s   0.1%  runtime.runSafePointFn
         0     0%  3.07%     -0.11s  0.55%  runtime.schedule
         0     0%  3.07%     -0.02s   0.1%  runtime.semacquire (inline)
         0     0%  3.07%     -0.02s   0.1%  runtime.semacquire1
         0     0%  3.07%     -0.14s   0.7%  runtime.semasleep
         0     0%  3.07%     -0.01s  0.05%  runtime.setprofilebucket
         0     0%  3.07%      0.06s   0.3%  runtime.slicebytetostring
         0     0%  3.07%     -0.05s  0.25%  runtime.startTheWorld
         0     0%  3.07%      0.23s  1.16%  runtime.startTheWorld.func1
         0     0%  3.07%      0.23s  1.16%  runtime.startTheWorldWithSema
         0     0%  3.07%      0.36s  1.81%  runtime.startm
         0     0%  3.07%      0.05s  0.25%  runtime.stealWork
         0     0%  3.07%     -0.15s  0.75%  runtime.stopTheWorld.func1
         0     0%  3.07%      0.04s   0.2%  runtime.stringtoslicebyte
         0     0%  3.07%      0.04s   0.2%  runtime.sysUnused (inline)
         0     0%  3.07%      0.04s   0.2%  runtime.sysUnusedOS
         0     0%  3.07%      0.02s   0.1%  runtime.sysUsed (inline)
         0     0%  3.07%      0.02s   0.1%  runtime.sysUsedOS
         0     0%  3.07%     -0.01s  0.05%  runtime.syscall_syscalln
         0     0%  3.07%     -0.01s  0.05%  runtime.typehash
         0     0%  3.07%     -0.02s   0.1%  runtime.unlock (partial-inline)
         0     0%  3.07%      0.07s  0.35%  runtime.unlock2Wake
         0     0%  3.07%     -0.02s   0.1%  runtime.unlockWithRank (inline)
         0     0%  3.07%      0.02s   0.1%  runtime.wakeNetPoller
         0     0%  3.07%      0.40s  2.01%  runtime.wakep
         0     0%  3.07%      0.01s  0.05%  runtime.wbBufFlush
         0     0%  3.07%      0.01s  0.05%  runtime.wbBufFlush.func1
         0     0%  3.07%      0.01s  0.05%  runtime.wbBufFlush1
         0     0%  3.07%     -0.02s   0.1%  strings.HasSuffix (inline)
         0     0%  3.07%     -0.04s   0.2%  strings.NewReader (inline)
         0     0%  3.07%      0.02s   0.1%  sync.(*Mutex).Lock (inline)
         0     0%  3.07%     -0.03s  0.15%  sync.(*Mutex).Unlock
         0     0%  3.07%     -0.01s  0.05%  sync.(*Once).Do (inline)
         0     0%  3.07%     -0.01s  0.05%  sync.(*Once).doSlow
         0     0%  3.07%     -0.01s  0.05%  sync.(*Pool).Put
         0     0%  3.07%     -0.01s  0.05%  syscall.Syscall9
         0     0%  3.07%      0.10s   0.5%  testing.(*B).StartTimer
         0     0%  3.07%     -0.49s  2.46%  testing.(*B).StopTimer
         0     0%  3.07%     -0.88s  4.42%  testing.(*B).launch
         0     0%  3.07%     -0.88s  4.42%  testing.(*B).runN
         0     0%  3.07%      0.01s  0.05%  testing.highPrecisionTimeNow (inline)
         0     0%  3.07%     -0.02s   0.1%  testing.highPrecisionTimeSince
         0     0%  3.07%     -0.01s  0.05%  testing.tRunner
         0     0%  3.07%     -0.01s  0.05%  time.(*Location).get
         0     0%  3.07%     -0.01s  0.05%  time.Time.Date
         0     0%  3.07%     -0.01s  0.05%  time.Time.absSec
         0     0%  3.07%     -0.01s  0.05%  time.abbrev
         0     0%  3.07%     -0.01s  0.05%  time.initLocal
         0     0%  3.07%     -0.01s  0.05%  time.initLocalFromTZI
         0     0%  3.07%     -0.01s  0.05%  time.matchZoneKey
         0     0%  3.07%     -0.01s  0.05%  time.toEnglishName
         0     0%  3.07%     -0.01s  0.05%  type:.eq.github.com/golang/mock/gomock.callSetKey