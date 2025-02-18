## Scenarios from OCIS API tests that are expected to fail with OCIS storage

The expected failures in this file are from features in the owncloud/ocis repo.

#### [Downloading the archive of the resource (files | folder) using resource path is not possible](https://github.com/owncloud/ocis/issues/4637)

- [apiArchiver/downloadByPath.feature:26](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiArchiver/downloadByPath.feature#L26)
- [apiArchiver/downloadByPath.feature:27](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiArchiver/downloadByPath.feature#L27)
- [apiArchiver/downloadByPath.feature:44](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiArchiver/downloadByPath.feature#L44)
- [apiArchiver/downloadByPath.feature:45](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiArchiver/downloadByPath.feature#L45)
- [apiArchiver/downloadByPath.feature:48](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiArchiver/downloadByPath.feature#L48)
- [apiArchiver/downloadByPath.feature:74](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiArchiver/downloadByPath.feature#L74)
- [apiArchiver/downloadByPath.feature:132](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiArchiver/downloadByPath.feature#L132)
- [apiArchiver/downloadByPath.feature:133](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiArchiver/downloadByPath.feature#L133)

### [Downloaded /Shares tar contains resource (files|folder) with leading / in Response](https://github.com/owncloud/ocis/issues/4636)

- [apiArchiver/downloadById.feature:134](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiArchiver/downloadById.feature#L134)
- [apiArchiver/downloadById.feature:135](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiArchiver/downloadById.feature#L135)

### [create request for already existing user exits with status code 500 ](https://github.com/owncloud/ocis/issues/3516)

- [apiGraph/createGroupCaseSensitive.feature:20](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/createGroupCaseSensitive.feature#L20)
- [apiGraph/createGroupCaseSensitive.feature:21](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/createGroupCaseSensitive.feature#L21)
- [apiGraph/createGroupCaseSensitive.feature:22](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/createGroupCaseSensitive.feature#L22)
- [apiGraph/createGroupCaseSensitive.feature:23](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/createGroupCaseSensitive.feature#L23)
- [apiGraph/createGroupCaseSensitive.feature:24](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/createGroupCaseSensitive.feature#L24)
- [apiGraph/createGroupCaseSensitive.feature:25](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/createGroupCaseSensitive.feature#L25)
- [apiGraph/createGroup.feature:28](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/createGroup.feature#L28)
- [apiGraph/createUser.feature:41](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/createUser.feature#L41)
- [apiGraph/createUser.feature:72](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/createUser.feature#L72)

### [PROPFIND on accepted shares with identical names containing brackets exit with 404](https://github.com/owncloud/ocis/issues/4421)

- [apiSpacesShares/changingFilesShare.feature:15](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiSpacesShares/changingFilesShare.feature#L15)

### [copy to overwrite (file and folder) from Personal to Shares Jail behaves differently](https://github.com/owncloud/ocis/issues/4393)

- [apiSpacesShares/copySpaces.feature:529](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiSpacesShares/copySpaces.feature#L529)
- [apiSpacesShares/copySpaces.feature:543](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiSpacesShares/copySpaces.feature#L543)

#### [PATCH request for TUS upload with wrong checksum gives incorrect response](https://github.com/owncloud/ocis/issues/1755)

- [apiSpacesShares/shareUploadTUS.feature:204](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiSpacesShares/shareUploadTUS.feature#L204)
- [apiSpacesShares/shareUploadTUS.feature:219](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiSpacesShares/shareUploadTUS.feature#L219)
- [apiSpacesShares/shareUploadTUS.feature:284](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiSpacesShares/shareUploadTUS.feature#L284)

### [Creating group with empty name returns status code 200](https://github.com/owncloud/ocis/issues/5050)

- [apiGraph/createGroup.feature:48](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/createGroup.feature#L48)

### [Settings service user can list other peoples assignments](https://github.com/owncloud/ocis/issues/5032)

- [apiAccountsHashDifficulty/assignRole.feature:28](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiAccountsHashDifficulty/assignRole.feature#L28)
- [apiAccountsHashDifficulty/assignRole.feature:29](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiAccountsHashDifficulty/assignRole.feature#L29)
- [apiGraph/assignRole.feature:31](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/assignRole.feature#L31)
- [apiGraph/assignRole.feature:32](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/assignRole.feature#L32)
- [apiGraph/assignRole.feature:33](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/assignRole.feature#L33)

#### [Share lists deleted user as 'user'](https://github.com/owncloud/ocis/issues/903)

- [apiGraph/deleteGroup.feature:68](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/deleteGroup.feature#L68)

#### [CORS headers are not identical with oC10 headers](https://github.com/owncloud/ocis/issues/5195)

- [apiCors/cors.feature:28](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiCors/cors.feature#L28)
- [apiCors/cors.feature:29](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiCors/cors.feature#L29)
- [apiCors/cors.feature:30](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiCors/cors.feature#L30)
- [apiCors/cors.feature:31](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiCors/cors.feature#L31)

#### [Requests with invalid credentials do not return CORS headers](https://github.com/owncloud/ocis/issues/5194)

- [apiCors/cors.feature:70](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiCors/cors.feature#L70)
- [apiCors/cors.feature:71](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiCors/cors.feature#L71)

#### [POST response does not return correct path when creating public link](https://github.com/owncloud/ocis/issues/5139)

- [apiSpacesShares/shareSubItemOfSpaceViaPublicLink.feature:63](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiSpacesShares/shareSubItemOfSpaceViaPublicLink.feature#L63)
- [apiSpacesShares/shareSubItemOfSpaceViaPublicLink.feature:64](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiSpacesShares/shareSubItemOfSpaceViaPublicLink.feature#L64)
- [apiSpacesShares/shareSubItemOfSpaceViaPublicLink.feature:65](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiSpacesShares/shareSubItemOfSpaceViaPublicLink.feature#L65)
- [apiSpacesShares/shareSubItemOfSpaceViaPublicLink.feature:93](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiSpacesShares/shareSubItemOfSpaceViaPublicLink.feature#L93)
- [apiSpacesShares/shareSubItemOfSpaceViaPublicLink.feature:167](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiSpacesShares/shareSubItemOfSpaceViaPublicLink.feature#L167)
- [apiSpacesShares/shareSubItemOfSpaceViaPublicLink.feature:168](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiSpacesShares/shareSubItemOfSpaceViaPublicLink.feature#L168)
- [apiSpacesShares/shareSubItemOfSpaceViaPublicLink.feature:169](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiSpacesShares/shareSubItemOfSpaceViaPublicLink.feature#L169)

#### [A User can get information of another user with Graph API](https://github.com/owncloud/ocis/issues/5125)

- [apiGraph/getUser.feature:83](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L83)
- [apiGraph/getUser.feature:84](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L84)
- [apiGraph/getUser.feature:85](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L85)
- [apiGraph/getUser.feature:86](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L86)
- [apiGraph/getUser.feature:87](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L87)
- [apiGraph/getUser.feature:88](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L88)
- [apiGraph/getUser.feature:89](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L89)
- [apiGraph/getUser.feature:90](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L90)
- [apiGraph/getUser.feature:91](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L91)
- [apiGraph/getUser.feature:92](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L92)
- [apiGraph/getUser.feature:93](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L93)
- [apiGraph/getUser.feature:94](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L94)
- [apiGraph/getUser.feature:607](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L607)
- [apiGraph/getUser.feature:608](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L608)
- [apiGraph/getUser.feature:609](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L609)
- [apiGraph/getUser.feature:610](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L610)
- [apiGraph/getUser.feature:611](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L611)
- [apiGraph/getUser.feature:612](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L612)
- [apiGraph/getUser.feature:613](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L613)
- [apiGraph/getUser.feature:614](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L614)
- [apiGraph/getUser.feature:615](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L615)
- [apiGraph/getUser.feature:616](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L616)
- [apiGraph/getUser.feature:617](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L617)
- [apiGraph/getUser.feature:618](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L618)

#### [Normal user can get expanded members information of a group](https://github.com/owncloud/ocis/issues/5604)

- [apiGraph/getGroup.feature:382](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getGroup.feature#L382)
- [apiGraph/getGroup.feature:383](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getGroup.feature#L383)
- [apiGraph/getGroup.feature:384](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getGroup.feature#L384)

#### [Changing user with an uppercase name gives 404 error](https://github.com/owncloud/ocis/issues/5763)

- [apiGraph/editUser.feature:68](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/editUser.feature#L68)

#### [Same users can be added in a group multiple time](https://github.com/owncloud/ocis/issues/5702)

- [apiGraph/addUserToGroup.feature:286](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L286)

#### [API requests from an unauthorized user should return 403](https://github.com/owncloud/ocis/issues/5938)

- [apiGraph/addUserToGroup.feature:151](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L151)
- [apiGraph/addUserToGroup.feature:152](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L152)
- [apiGraph/addUserToGroup.feature:153](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L153)
- [apiGraph/addUserToGroup.feature:185](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L185)
- [apiGraph/addUserToGroup.feature:186](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L186)
- [apiGraph/addUserToGroup.feature:187](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L187)
- [apiGraph/createGroup.feature:43](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/createGroup.feature#L43)
- [apiGraph/createGroup.feature:44](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/createGroup.feature#L44)
- [apiGraph/createGroup.feature:45](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/createGroup.feature#L45)
- [apiGraph/deleteGroup.feature:63](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/deleteGroup.feature#L63)
- [apiGraph/deleteGroup.feature:64](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/deleteGroup.feature#L64)
- [apiGraph/deleteGroup.feature:65](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/deleteGroup.feature#L65)
- [apiGraph/editGroup.feature:35](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/editGroup.feature#L35)
- [apiGraph/editGroup.feature:36](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/editGroup.feature#L36)
- [apiGraph/editGroup.feature:37](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/editGroup.feature#L37)
- [apiGraph/getGroup.feature:55](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getGroup.feature#L55)
- [apiGraph/getGroup.feature:56](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getGroup.feature#L56)
- [apiGraph/getGroup.feature:57](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getGroup.feature#L57)
- [apiGraph/getGroup.feature:104](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getGroup.feature#L104)
- [apiGraph/getGroup.feature:105](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getGroup.feature#L105)
- [apiGraph/getGroup.feature:106](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getGroup.feature#L106)
- [apiGraph/getGroup.feature:268](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getGroup.feature#L268)
- [apiGraph/getGroup.feature:269](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getGroup.feature#L269)
- [apiGraph/getGroup.feature:270](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getGroup.feature#L270)
- [apiGraph/removeUserFromGroup.feature:192](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/removeUserFromGroup.feature#L192)
- [apiGraph/removeUserFromGroup.feature:193](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/removeUserFromGroup.feature#L193)
- [apiGraph/removeUserFromGroup.feature:194](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/removeUserFromGroup.feature#L194)

#### [API requests for a non-existent resources should return 404](https://github.com/owncloud/ocis/issues/5939)

- [apiGraph/addUserToGroup.feature:202](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L202)
- [apiGraph/addUserToGroup.feature:203](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L203)
- [apiGraph/addUserToGroup.feature:204](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L204)

### [Users are added in a group with wrong host in host-part of user](https://github.com/owncloud/ocis/issues/5871)

- [apiGraph/addUserToGroup.feature:370](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L370)
- [apiGraph/addUserToGroup.feature:384](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L384)

### [Adding the same user as multiple members in a single request results in listing the same user twice in the group](https://github.com/owncloud/ocis/issues/5855)

- [apiGraph/addUserToGroup.feature:421](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L421)

Note: always have an empty line at the end of this file.
The bash script that processes this file requires that the last line has a newline on the end.
