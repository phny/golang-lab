func preCluster() error {
        file, err := os.OpenFile(*featureFile, os.O_RDWR, 0666)
        if err != nil {
                log.Fatal("faled to open face feature file, err: ", err)
        }
        defer file.Close()

        featContent, err := ioutil.ReadAll(file)
        if err != nil {
                log.Fatal("failed to read face feature file, err: ", err)
        }
        featLen := len(featContent) / featureLen
        log.Infof("featLen:%d", featLen)
        featIds := make([]int64, 0)

        for i := 0; i < 32; i++ {
                featIds = append(featIds, int64(i+10000))
        }

        // accumulate data
        var data []byte
        fdata := make([]float32, 0)
        data = featContent[*begin*featureLen*32 : *end*featureLen*32]
        i := 0
        for i < 32 {
                ff := utils.EncodeFeatureByte2Float(data[i*featureLen : (i+1)*featureLen])
                log.Infof("ff: %v", ff)
                fdata = append(fdata, ff...)
                i++
        }

        log.Infof("fdata: %d", len(fdata))
        if err := sego.SearchEnigineInit("Kestrel", licPath); err != nil {
                log.Errorf("search engine init failed: %v", err)
                return err
        }
        defer sego.SearchEngineDeinit()
        sego.SearchEngineSetLogLevel(sego.SELogLevelTrace)

        ctx, err := sego.CreateContext(0)
        if err != nil {
                log.Errorf("creat context failed: %v", err)
                return err
        }
        defer ctx.Destroy()

        flatIndex, err := sego.CreateFlatIndex(ctx, featureDim, useInt8)
        if err != nil {
                log.Errorf("create flat index failed: %v", err)
                return err
        }
        defer flatIndex.Destroy()

        featAggr, err := sego.NewFeatureAggregator(*flatIndex, 2, 32, 0.5, 0.3, 1, 0)
        if err != nil {
                log.Errorf("new feature aggregator failed: %v", err)
                return err
        }
        defer featAggr.Destroy()

        agrrFeatures, err := featAggr.Aggregate(32, fdata)
        if err != nil {
                log.Fatal(err)
        }

        dbscanCluster, err := sego.NewDbsacnClusterer(*flatIndex, 0.9, 1)
        if err != nil {
                log.Fatalf("failed to init dbscan clusterer, err: %v", err)
        }
        log.Infof("featIds: %d", len(featIds))
        labels, err := dbscanCluster.Clustering(32, featIds, agrrFeatures)
        if err != nil {
                log.Fatalf("feature cluster failed, err: %v", err)
        }
        log.Infof("labels:%v", labels)
        length := realtime.NoRepeatedNumInUnorderedArray(labels)
        vec := make([][]int32, length+1)
        for i, label := range labels {
                vec[label] = append(vec[label], int32(i))
        }
        log.Infof("vec: %v", vec)
        return nil
}
