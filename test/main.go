package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)
type JSONData struct {
	Devices []struct {
		ChangeSectorCnt int `json:"change_sector_cnt"`
		DcrSectorBegin int `json:"dcr_sector_begin"`
		DcrSectorEnd int `json:"dcr_sector_end"`
		DcrSectorSnd int `json:"dcr_sector_snd"`
		DdrSectorBegin int `json:"ddr_sector_begin"`
		DdrSectorEnd int `json:"ddr_sector_end"`
		DdrSectorSeqLast int `json:"ddr_sector_seq_last"`
		DdrSectorSnd int `json:"ddr_sector_snd"`
		DdrTsEnd interface{} `json:"ddr_ts_end"`
		DdrTsLaunch interface{} `json:"ddr_ts_launch"`
		Enable int `json:"enable"`
		Hole interface{} `json:"hole"`
		NsdmSectorSnd int `json:"nsdm_sector_snd"`
		NsdmSetCnt int `json:"nsdm_set_cnt"`
		OrName string `json:"or_name"`
		PIdrMp int64 `json:"p_idr_mp"`
		Role int `json:"role"`
		SectorCntOfMdHead int `json:"sector_cnt_of_md_head"`
		SectorSeqLast int `json:"sector_seq_last"`
		SectorSizeMod int `json:"sector_size_mod"`
		SectorSizePowerNum int `json:"sector_size_power_num"`
		Version string `json:"version"`
	} `json:"devices"`
	Groups []struct {
		LocalIP string `json:"LOCAL_IP"`
		LocalLinkID string `json:"LOCAL_LINK_ID"`
		LocalPort string `json:"LOCAL_PORT"`
		LogLevel string `json:"LOG_LEVEL"`
		PeerIP string `json:"PEER_IP"`
		PeerPort string `json:"PEER_PORT"`
		Async int `json:"async"`
		CdpDisabledByReplicating int `json:"cdp_disabled_by_replicating"`
		CdpDisabledByRetrieving int `json:"cdp_disabled_by_retrieving"`
		CdpDisabledByRolling int `json:"cdp_disabled_by_rolling"`
		CipherEnb int `json:"cipher_enb"`
		CipherKlen int `json:"cipher_klen"`
		ComprehensiveGrpState string `json:"comprehensive_grp_state"`
		CompressEnb int `json:"compress_enb"`
		Connection int `json:"connection"`
		DataStampNsec int `json:"data_stamp_nsec"`
		DataStampSec int `json:"data_stamp_sec"`
		DevRef int `json:"dev_ref"`
		DisableSelfAdapt int `json:"disable_self_adapt"`
		EnableAppWriteComplete int `json:"enable_app_write_complete"`
		EnableIoOrderChk int `json:"enable_io_order_chk"`
		EnableLogInStg int `json:"enable_log_in_stg"`
		EnableReuseaddr int `json:"enable_reuseaddr"`
		Enabled int `json:"enabled"`
		Errcnt int `json:"errcnt"`
		GrpName string `json:"grp_name"`
		HashAlgorithmName int `json:"hash_algorithm_name"`
		HaveDcr int `json:"have_dcr"`
		HaveNsdm int `json:"have_nsdm"`
		HaveOrl int `json:"have_orl"`
		HaveOrls int `json:"have_orls"`
		HeadTimeNsec int `json:"head_time_nsec"`
		HeadTimeSec int `json:"head_time_sec"`
		Inconsistent int `json:"inconsistent"`
		Integral int `json:"integral"`
		LocalStamp int64 `json:"local_stamp"`
		LogCnt int `json:"log_cnt"`
		LogCntHold int `json:"log_cnt_hold"`
		LogCntPowerNum int `json:"log_cnt_power_num"`
		LogIdxLpIdxHead int `json:"log_idx_lp_idx_head"`
		LogIdxLpIdxTail int `json:"log_idx_lp_idx_tail"`
		LogIdxPageCnt int `json:"log_idx_page_cnt"`
		LogIdxPageHead int `json:"log_idx_page_head"`
		LogIdxPageTail int `json:"log_idx_page_tail"`
		NodeCntPowerNum int `json:"node_cnt_power_num"`
		OriginUTC string `json:"originUTC"`
		PauseByLocal int `json:"pause_by_local"`
		PauseByPeer int `json:"pause_by_peer"`
		PeerStamp int `json:"peer_stamp"`
		RateAckByte int `json:"rate_ack_byte"`
		RateSmtByte int `json:"rate_smt_byte"`
		ReadChkPoint int `json:"read_chk_point"`
		Role int `json:"role"`
		SectorCntOfLogIdxOffset int `json:"sector_cnt_of_log_idx_offset"`
		SectorCntOfMdHead int `json:"sector_cnt_of_md_head"`
		SectorCntOfTotalUseMd int `json:"sector_cnt_of_total_use_md"`
		SktDataByUDP int `json:"skt_data_by_udp"`
		SndPktPowerNum int `json:"snd_pkt_power_num"`
		SndWindow int `json:"snd_window"`
		TailTimeNsec int `json:"tail_time_nsec"`
		TailTimeSec int `json:"tail_time_sec"`
		TCPState int `json:"tcp_state"`
		ThresholdPend int `json:"threshold_pend"`
		ThresholdTruncate int `json:"threshold_truncate"`
		ThresholdTruncateMargin int `json:"threshold_truncate_margin"`
		UnusedOrl int `json:"unused_orl"`
		UpToDate int `json:"up_to_date"`
		Validity int `json:"validity"`
		ValidityDays int64 `json:"validity_days"`
		Version string `json:"version"`
		Warncnt int `json:"warncnt"`
	} `json:"groups"`
}

func main() {
	cmd := exec.Command("/usr/local/GoCode/shadow/bin/mac/Sclient","EXEC","-Rserver","10.200.20.253:3721","-user","root", "-cmd", "/etc/or.d/orioctl get orbd")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
    	fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
    	return
	}
	var jsonData JSONData
	err = json.Unmarshal(out.Bytes(),&jsonData)
	if err != nil {
		log.Fatal("error \n",err)
	}
	fmt.Printf("%#v",jsonData.Devices[0])

}
