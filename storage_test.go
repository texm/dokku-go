package dokku

func (s *DokkuTestSuite) TestManageStorage() {
	r := s.Require()

	appName := "test-storage-app"

	r.NoError(s.Client.CreateApp(appName))

	storageReport, err := s.Client.GetAppStorageReport(appName)
	r.NoError(err)
	r.Len(storageReport.RunMounts, 0)

	storage := StorageBindMount{
		HostDir:      "testAppStorage",
		ContainerDir: "/data",
	}
	r.NoError(s.Client.EnsureStorageDirectory(storage.HostDir, StorageChownOptionHerokuish))
	r.NoError(s.Client.MountAppStorage(appName, storage))

	storageList, err := s.Client.ListAppStorage(appName)
	r.NoError(err)
	r.Len(storageList, 1)
	r.Equal(storage, storageList[0])
}
