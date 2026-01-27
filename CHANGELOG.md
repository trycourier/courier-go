# Changelog

## 4.5.0 (2026-01-27)

Full Changelog: [v4.4.3...v4.5.0](https://github.com/trycourier/courier-go/compare/v4.4.3...v4.5.0)

### Features

* **client:** add a convenient param.SetJSON helper ([4666d08](https://github.com/trycourier/courier-go/commit/4666d08f19c1a07c74dddf2ed2aa18db01686ef7))


### Bug Fixes

* **docs:** add missing pointer prefix to api.md return types ([f279e07](https://github.com/trycourier/courier-go/commit/f279e070072a6304c14a007e6ca680a0946f68c3))


### Chores

* **internal:** update `actions/checkout` version ([95dba85](https://github.com/trycourier/courier-go/commit/95dba85c35d0a539a67e063acd776b20e767d4d7))


### Documentation

* **types:** clarify version field description in AutomationTemplate ([1711973](https://github.com/trycourier/courier-go/commit/17119733aa5efb4325b224c9e99b5ec4bbaa8f89))

## 4.4.3 (2026-01-14)

Full Changelog: [v4.4.2...v4.4.3](https://github.com/trycourier/courier-go/compare/v4.4.2...v4.4.3)

### Chores

* **internal:** regenerate SDK with no functional changes ([5ec8bb4](https://github.com/trycourier/courier-go/commit/5ec8bb41516427f5538c4b2f1c83d4129fc366ac))
* **internal:** regenerate SDK with no functional changes ([8cd6cd1](https://github.com/trycourier/courier-go/commit/8cd6cd1f34fbc356fc6e0369e47870df4167573d))
* remove custom code ([931289a](https://github.com/trycourier/courier-go/commit/931289a85bbb08f9a6d852364536cd71bebba8e5))

## 4.4.2 (2026-01-12)

Full Changelog: [v4.4.1...v4.4.2](https://github.com/trycourier/courier-go/compare/v4.4.1...v4.4.2)

### Bug Fixes

* **types:** remove FilterConfigUnion type, consolidate to FilterUnion in audience ([f7eb351](https://github.com/trycourier/courier-go/commit/f7eb35124e3bb1e323c07550951a3d49355ceb1c))


### Chores

* **internal:** regenerate SDK with no functional changes ([1a2515a](https://github.com/trycourier/courier-go/commit/1a2515ae6f8d35d883bc870d31dce248c41cce5f))

## 4.4.1 (2026-01-08)

Full Changelog: [v4.4.0...v4.4.1](https://github.com/trycourier/courier-go/compare/v4.4.0...v4.4.1)

### Bug Fixes

* **api:** remove audit-events, automations, brands, bulk, inbound, translations endpoints ([3f87200](https://github.com/trycourier/courier-go/commit/3f8720082e80b4b06bdb4392d8f4509d2591d2e1))


### Chores

* **internal:** codegen related update ([af9d61c](https://github.com/trycourier/courier-go/commit/af9d61c692565e2b449f41dc26440e3bab848064))

## 4.4.0 (2025-12-22)

Full Changelog: [v4.3.0...v4.4.0](https://github.com/trycourier/courier-go/compare/v4.3.0...v4.4.0)

### Features

* **api:** add audience/slack/msteams/pagerduty/webhook recipients, refactor to union types ([353c0ac](https://github.com/trycourier/courier-go/commit/353c0ac79473b3913e21d74e2a9eabdcbfe40cf3))


### Bug Fixes

* **client:** properly marshal embedded structs ([75c8698](https://github.com/trycourier/courier-go/commit/75c86984ec2c1bea7ab85ecda29b7b31ee920bde))
* skip usage tests that don't work with Prism ([b8bce4c](https://github.com/trycourier/courier-go/commit/b8bce4c295a9af0852092557d68d40f9390c324c))


### Chores

* add float64 to valid types for RegisterFieldValidator ([82c9488](https://github.com/trycourier/courier-go/commit/82c94885b40fcf49be0970c91b33c5889a0333bd))

## 4.3.0 (2025-12-16)

Full Changelog: [v4.2.1...v4.3.0](https://github.com/trycourier/courier-go/compare/v4.2.1...v4.3.0)

### Features

* Add timezone field to Delay schema ([82022e4](https://github.com/trycourier/courier-go/commit/82022e483e3944ca50e3f03e8bae3c71bb25259f))
* **encoder:** support bracket encoding form-data object members ([9e47be4](https://github.com/trycourier/courier-go/commit/9e47be4d14f2a70505e878c8fd9d3dd176edeb28))
* Fix UsersGetAllTokensResponse to return object with tokens property i… ([8f24b6d](https://github.com/trycourier/courier-go/commit/8f24b6d0f07c4621c8f0cb3df536f9424dd1d760))
* Update bulk API spec: make event required, document profile.email req… ([7c385f3](https://github.com/trycourier/courier-go/commit/7c385f3179886c9eab19653f23946b7ac898cfe2))

## 4.2.1 (2025-12-08)

Full Changelog: [v4.2.0...v4.2.1](https://github.com/trycourier/courier-go/compare/v4.2.0...v4.2.1)

## 4.2.0 (2025-12-08)

Full Changelog: [v4.1.1...v4.2.0](https://github.com/trycourier/courier-go/compare/v4.1.1...v4.2.0)

### Features

* Add event_ids field to Notification schema ([ea63268](https://github.com/trycourier/courier-go/commit/ea63268e6175851811f3da32b8ed6ab4c8f6b1eb))


### Bug Fixes

* **client:** fix duplicate Go struct resulting from name derivations schema ([dc40006](https://github.com/trycourier/courier-go/commit/dc40006ded7f953c89e1b23b5e917bc986be33f5))
* **mcp:** correct code tool API endpoint ([e3f20a5](https://github.com/trycourier/courier-go/commit/e3f20a5e9922b18f49728849810d11e4cd57c57e))
* rename param to avoid collision ([9ac4b56](https://github.com/trycourier/courier-go/commit/9ac4b5691091dc59381250bfff644b2c3c9f1dd0))


### Chores

* elide duplicate aliases ([5f4976c](https://github.com/trycourier/courier-go/commit/5f4976cedee15db7a4713944c20fc07de8c86bb2))
* **internal:** codegen related update ([64c0731](https://github.com/trycourier/courier-go/commit/64c073156678a16f20821303815eab0e26568bdf))

## 4.1.1 (2025-12-02)

Full Changelog: [v4.1.0...v4.1.1](https://github.com/trycourier/courier-go/compare/v4.1.0...v4.1.1)

### Bug Fixes

* **client:** correctly specify Accept header with */* instead of empty ([e648abc](https://github.com/trycourier/courier-go/commit/e648abc00d93fc310cb287ea2d0188f2609bec49))


### Chores

* fix empty interfaces ([47b7ec0](https://github.com/trycourier/courier-go/commit/47b7ec0afa5f7e63938c045074aedb04ce442e02))

## 4.1.0 (2025-11-18)

Full Changelog: [v4.0.0...v4.1.0](https://github.com/trycourier/courier-go/compare/v4.0.0...v4.1.0)

### Features

* JWT scope updates ([ec65b4d](https://github.com/trycourier/courier-go/commit/ec65b4d201c05db7dd79386a673d7840113c9276))
* NPM enabled ([f74d261](https://github.com/trycourier/courier-go/commit/f74d26190ca244212821dcc8dca481e32489cfd8))
* Test ([74a9882](https://github.com/trycourier/courier-go/commit/74a9882eab3e90ad2957bd2fb4eb21813d66ac91))

## 4.0.0 (2025-11-12)

Full Changelog: [v3.3.0-alpha7...v4.0.0](https://github.com/trycourier/courier-go/compare/v3.3.0-alpha7...v4.0.0)

### Features

* Spec Comment Change ([3cfc868](https://github.com/trycourier/courier-go/commit/3cfc868dc6627594a89ab480b68d17140d4bed60))
* Token Prop Description Change ([674a316](https://github.com/trycourier/courier-go/commit/674a316fc5a2f8e3ab9b6bf4e110f5bc5dcbcf88))


### Chores

* bump gjson version ([13a17be](https://github.com/trycourier/courier-go/commit/13a17be31763e8e8338fd7d029dcfe69e2fbb181))
* update SDK settings ([92b0a3f](https://github.com/trycourier/courier-go/commit/92b0a3fd4d44bd2b2bfb15b1fd4a16fd9724e9ac))

## 3.3.0-alpha7 (2025-11-07)

Full Changelog: [v3.3.0-alpha6...v3.3.0-alpha7](https://github.com/trycourier/courier-go/compare/v3.3.0-alpha6...v3.3.0-alpha7)

### Features

* Attempt kick off again ([7cf87be](https://github.com/trycourier/courier-go/commit/7cf87be0a3dc83e62b8accac199862b3527ea5cd))
* Organization update ([3231738](https://github.com/trycourier/courier-go/commit/3231738fa25b6404c56db14e76076db694025b59))


### Bug Fixes

* Better Python Samples + Updates to naming ([7da5eaa](https://github.com/trycourier/courier-go/commit/7da5eaa269a957d2bf6ff490f1ac732fde02e9f8))


### Chores

* **internal:** grammar fix (it's -&gt; its) ([1fb1631](https://github.com/trycourier/courier-go/commit/1fb16319bc54b410b5a7edabf2dbe148a8b11ceb))

## 3.3.0-alpha6 (2025-10-31)

Full Changelog: [v3.3.0-alpha5...v3.3.0-alpha6](https://github.com/trycourier/courier-go/compare/v3.3.0-alpha5...v3.3.0-alpha6)

### Features

* Comment adjustment to kick of build ([726bb89](https://github.com/trycourier/courier-go/commit/726bb89467326b0272c3e0b2e7045b9b5d081ccd))


### Bug Fixes

* Comment to kick off build ([67986ea](https://github.com/trycourier/courier-go/commit/67986ea8a06a7ad022e7983c6e9976831d8b3731))

## 3.3.0-alpha5 (2025-10-17)

Full Changelog: [v3.3.0-alpha4...v3.3.0-alpha5](https://github.com/trycourier/courier-go/compare/v3.3.0-alpha4...v3.3.0-alpha5)

### Bug Fixes

* Dep Warning ([237e89a](https://github.com/trycourier/courier-go/commit/237e89a01569652d2948585bfafa78b3a934a326))

## 3.3.0-alpha4 (2025-10-17)

Full Changelog: [v3.3.0-alpha3...v3.3.0-alpha4](https://github.com/trycourier/courier-go/compare/v3.3.0-alpha3...v3.3.0-alpha4)

### Bug Fixes

* Updated paths for each model and go example updates ([6a2f8ce](https://github.com/trycourier/courier-go/commit/6a2f8ce17540ac71d737d1e2fa44761c3f35d0e3))

## 3.3.0-alpha3 (2025-10-17)

Full Changelog: [v3.3.0-alpha2...v3.3.0-alpha3](https://github.com/trycourier/courier-go/compare/v3.3.0-alpha2...v3.3.0-alpha3)

### Features

* Changes to spec, examples and scripts ([ce585a5](https://github.com/trycourier/courier-go/commit/ce585a5714c7101ce074d07b8fb631c2ba27c3ff))
* More PHP and attempted node automerge ([1a423ee](https://github.com/trycourier/courier-go/commit/1a423eee87e690b33bd8313a95c10ea846e37f16))


### Bug Fixes

* **tests:** use correct param type for preferences ([fc89e7a](https://github.com/trycourier/courier-go/commit/fc89e7a3af645c464055b556140dfd0c43704bbc))

## 3.3.0-alpha2 (2025-10-14)

Full Changelog: [v3.3.0-alpha1...v3.3.0-alpha2](https://github.com/trycourier/courier-go/compare/v3.3.0-alpha1...v3.3.0-alpha2)

### Features

* Kick of merge attempt ([c47fed3](https://github.com/trycourier/courier-go/commit/c47fed302f24119b7a8d4f444bf0c96fc9f11860))
* Polish and Kick of Java Kit Gen ([63588bc](https://github.com/trycourier/courier-go/commit/63588bc4ce5f8845ddc71656890899a20deafa4b))

## 3.3.0-alpha1 (2025-10-09)

Full Changelog: [v3.3.0-alpha0...v3.3.0-alpha1](https://github.com/trycourier/courier-go/compare/v3.3.0-alpha0...v3.3.0-alpha1)

### Features

* Template Id ([d1e43b8](https://github.com/trycourier/courier-go/commit/d1e43b80564e7815cf3379528426f09d04fc3b27))

## 3.3.0-alpha0 (2025-10-09)

Full Changelog: [v0.0.1...v3.3.0-alpha0](https://github.com/trycourier/courier-go/compare/v0.0.1...v3.3.0-alpha0)

### Features

* **api:** manual updates ([9ee30fb](https://github.com/trycourier/courier-go/commit/9ee30fbd3e2124447b2531dd5c41481ce51a5fbd))
* **api:** manual updates ([29c5130](https://github.com/trycourier/courier-go/commit/29c5130c84418466ab3b2b4161e06eca3236f101))
* **api:** manual updates ([996c0e7](https://github.com/trycourier/courier-go/commit/996c0e71a9f96056e4e23329a683a43a5db36142))
* **api:** manual updates ([8bce648](https://github.com/trycourier/courier-go/commit/8bce648faf7704b75dc0e114eb0a8b04e89a70ef))
* **api:** manual updates ([1952774](https://github.com/trycourier/courier-go/commit/19527742d2c591da5fb09cd917df8d5d9ea66dee))
* **api:** manual updates ([692331e](https://github.com/trycourier/courier-go/commit/692331efbab3b5f0aed24464d6099c453d9ee66e))
* **api:** manual updates ([22051c9](https://github.com/trycourier/courier-go/commit/22051c9b249d43167b9f93ad597d39359f6c05df))
* **api:** manual updates ([b41a673](https://github.com/trycourier/courier-go/commit/b41a673d3096675949aec38acb147d039c726cc4))
* **api:** manual updates ([d39adc4](https://github.com/trycourier/courier-go/commit/d39adc4b6c09138d88d84335e339524e22961cb1))
* **api:** manual updates ([d6ef959](https://github.com/trycourier/courier-go/commit/d6ef959bd3c20bf5ef3fc0bf0dfe062b2b30cf5a))
* **api:** manual updates ([a403dff](https://github.com/trycourier/courier-go/commit/a403dffdef3d13b992987d227bd5f4381fdf1ced))
* **api:** manual updates ([6c16c65](https://github.com/trycourier/courier-go/commit/6c16c6593a8dc589429c1dc48c0d28a34009ee87))
* **api:** manual updates ([c8ad6ca](https://github.com/trycourier/courier-go/commit/c8ad6cac2899e758788ac911596ae0bcdb0b7db6))
* **api:** manual updates ([e838734](https://github.com/trycourier/courier-go/commit/e838734305b5a3b81045c0673b8faa4dca55948c))
* **api:** manual updates ([a0cd443](https://github.com/trycourier/courier-go/commit/a0cd443a1e750b94dbf32a54931c3ae25f118a2d))
* **api:** manual updates ([2fe0cb9](https://github.com/trycourier/courier-go/commit/2fe0cb93997ea8acf03455242824c2a3c0e49ebe))
* **api:** manual updates ([e5b78d9](https://github.com/trycourier/courier-go/commit/e5b78d94ee0ec5df38444d4581402ed1bdd55314))
* **api:** manual updates ([7207a3d](https://github.com/trycourier/courier-go/commit/7207a3dba035a037ea3c3110b6fae57f12a769f2))
* **api:** manual updates ([38caab0](https://github.com/trycourier/courier-go/commit/38caab0b0218cfd5dc578a871ca9f41ad997356d))
* **api:** manual updates ([2653e7d](https://github.com/trycourier/courier-go/commit/2653e7dfb4cbfa6d163c0686dc445851562bf7bb))
* **api:** manual updates ([8b6bb20](https://github.com/trycourier/courier-go/commit/8b6bb2018b5dfc6862d0acbf73d8c95cfe8e102e))
* **api:** manual updates ([cfe7ef3](https://github.com/trycourier/courier-go/commit/cfe7ef3735a4927c4e07b0036707db6f9935e375))
* **api:** manual updates ([9cd1250](https://github.com/trycourier/courier-go/commit/9cd1250ea38a8b4315f7e78bd842a6def3d50941))
* **api:** manual updates ([c84a4a0](https://github.com/trycourier/courier-go/commit/c84a4a00074fe788ead10a376b4b6fa8ab8f5175))
* **api:** manual updates ([23c1247](https://github.com/trycourier/courier-go/commit/23c12470f36a41e6add7af02db9a9c92a70c5e60))
* **api:** manual updates ([c28c1f5](https://github.com/trycourier/courier-go/commit/c28c1f5dcdca98b276fa28652238bceeee599b17))
* **api:** manual updates ([04122e8](https://github.com/trycourier/courier-go/commit/04122e89317218750c9a2a13ee2e47963f401340))
* **api:** manual updates ([073c81d](https://github.com/trycourier/courier-go/commit/073c81dadd8f222264ffe1e8e35adb6fc0ae566a))
* **api:** manual updates ([6a7e395](https://github.com/trycourier/courier-go/commit/6a7e395047082175229fbebf18552d7c1adbf066))
* **api:** manual updates ([09aa3c0](https://github.com/trycourier/courier-go/commit/09aa3c09972673362a12d2467dc0909dbbe853f3))
* **api:** manual updates ([8f7e061](https://github.com/trycourier/courier-go/commit/8f7e0618ec1cf431b521c425b8dae22db86ba169))
* Examples and ref polish ([be6b04e](https://github.com/trycourier/courier-go/commit/be6b04efbbae3ca01e3cb9532f3af7661acbc2d5))
* Model sync ([9b5ec45](https://github.com/trycourier/courier-go/commit/9b5ec45331d89ee5f33121782f23ee962253f805))
* Test Github Action ([a13ca2d](https://github.com/trycourier/courier-go/commit/a13ca2d97b8464dcf844df517e4bd3e5374154ee))


### Chores

* configure new SDK language ([6113aa5](https://github.com/trycourier/courier-go/commit/6113aa51e63bf0ec71898f90b42f8332fe84a22c))
* update SDK settings ([31c2cf9](https://github.com/trycourier/courier-go/commit/31c2cf91cf37761c025e08287496326986d17dcc))
* update SDK settings ([313f6a0](https://github.com/trycourier/courier-go/commit/313f6a08c6d9dce293177b90f3749b6a045af8aa))
