// Copyright (c) 2022 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

#[cfg(target_arch = "x86_64")]
pub use arch_specific::*;

mod arch_specific {
    use crate::check;
    use anyhow::{anyhow, Result};

    const PROC_CPUINFO: &str = "/proc/cpuinfo";
    const CPUINFO_DELIMITER: &str = "\nprocessor";
    const CPUINFO_FLAGS_TAG: &str = "flags";
    const CPU_FLAGS_INTEL: &[&str] = &["lm", "sse4_1", "vmx"];
    const CPU_ATTRIBS_INTEL: &[&str] = &["GenuineIntel"];

    // check cpu
    fn check_cpu() -> Result<()> {
        println!("INFO: check CPU: x86_64");

        let cpu_info = check::get_single_cpu_info(PROC_CPUINFO, CPUINFO_DELIMITER)?;

        let cpu_flags = check::get_cpu_flags(&cpu_info, CPUINFO_FLAGS_TAG)
            .map_err(|e| anyhow!("Error parsing CPU flags, file {:?}, {:?}", PROC_CPUINFO, e))?;

        // perform checks
        // TODO: Perform checks based on hypervisor type
        // TODO: Add more information to output (see kata-check in go tool); adjust formatting
        let missing_cpu_attributes = check::check_cpu_attribs(&cpu_info, CPU_ATTRIBS_INTEL)?;
        if !missing_cpu_attributes.is_empty() {
            eprintln!(
                "WARNING: Missing CPU attributes {:?}",
                missing_cpu_attributes
            );
        }
        let missing_cpu_flags = check::check_cpu_flags(&cpu_flags, CPU_FLAGS_INTEL)?;
        if !missing_cpu_flags.is_empty() {
            eprintln!("WARNING: Missing CPU flags {:?}", missing_cpu_flags);
        }

        Ok(())
    }

    pub fn check() -> Result<()> {
        println!("INFO: check: x86_64");

        let _cpu_result = check_cpu();

        // TODO: collect outcome of tests to determine if checks pass or not

        Ok(())
    }
}
