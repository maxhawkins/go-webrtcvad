#ifndef _WIN32
	#define WEBRTC_POSIX
#endif

#include "webrtc/common_audio/signal_processing/complex_bit_reverse.c"
#include "webrtc/common_audio/signal_processing/complex_fft.c"
#include "webrtc/common_audio/signal_processing/cross_correlation.c"
#include "webrtc/common_audio/signal_processing/division_operations.c"
#include "webrtc/common_audio/signal_processing/downsample_fast.c"
#include "webrtc/common_audio/signal_processing/energy.c"
#include "webrtc/common_audio/signal_processing/get_scaling_square.c"
#include "webrtc/common_audio/signal_processing/min_max_operations.c"
#include "webrtc/common_audio/signal_processing/real_fft.c"
#include "webrtc/common_audio/signal_processing/resample_48khz.c"
#include "webrtc/common_audio/signal_processing/resample_by_2_internal.c"
#include "webrtc/common_audio/signal_processing/resample_fractional.c"
#include "webrtc/common_audio/signal_processing/spl_init.c"
#include "webrtc/common_audio/signal_processing/vector_scaling_operations.c"
#include "webrtc/common_audio/vad/vad_core.c"
#include "webrtc/common_audio/vad/vad_filterbank.c"
#include "webrtc/common_audio/vad/vad_gmm.c"
#include "webrtc/common_audio/vad/vad_sp.c"
#include "webrtc/common_audio/vad/webrtc_vad.c"
